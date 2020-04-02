package runner

import (
	"context"

	"github.com/jonyhy96/lgtm-action/pkg/util/env"
	githubutil "github.com/jonyhy96/lgtm-action/pkg/util/github"
	"github.com/jonyhy96/lgtm-action/pkg/util/owner"

	"github.com/google/go-github/v30/github"
	"github.com/pkg/errors"
)

var (
	approve = "APPROVE"
)

var (
	// ErrorCantGetNumber happens when get number error.
	ErrorCantGetNumber = errors.New("can't get number from GITHUB_EVENT_PATH")
	// ErrorNotEnoughTimes happens when approve time is less than given arg times.
	ErrorNotEnoughTimes = errors.New("not enough times of approve")
)

// Runner entity.
type Runner struct {
	client *github.Client
	ctx    *context.Context
}

// New creates a new Runner.
func New(ctx *context.Context, client *github.Client) Runner {
	return Runner{
		client: client,
		ctx:    ctx,
	}
}

// Run run lgtm logic.
func (r *Runner) Run(times int, ownersFile string) error {
	owners, err := owner.GetALL(ownersFile)
	if err != nil {
		return err
	}
	var ownersMap = make(map[string]interface{})
	for _, owner := range owners {
		ownersMap[owner] = struct{}{}
	}

	event, err := githubutil.GetEvent()
	if err != nil {
		return err
	}

	var number = event.Issue.Number
	if event.PullRequest.Number != 0 {
		number = event.PullRequest.Number
	}
	if number == 0 {
		return ErrorCantGetNumber
	}

	comments, _, err := r.client.Issues.ListComments(*r.ctx, env.Owner, env.Repository, number, nil)
	if err != nil {
		return err
	}

	var approverMap = make(map[string]interface{})
	for _, comment := range comments {
		login := comment.User.GetLogin()
		if _, ok := ownersMap[login]; ok {
			approverMap[login] = struct{}{}
		}
	}

	if len(approverMap) < times {
		return ErrorNotEnoughTimes
	}

	_, _, err = r.client.PullRequests.CreateReview(*r.ctx, env.Owner, env.Repository, number, &github.PullRequestReviewRequest{
		Event: &approve,
	})
	if err != nil {
		return err
	}

	return nil
}
