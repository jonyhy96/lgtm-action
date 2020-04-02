package github

import (
	"context"
	"encoding/json"
	"os"

	"github.com/google/go-github/v30/github"
	"golang.org/x/oauth2"
)

// Event entity.
type Event struct {
	Issue       Issue       `json:"issue"`
	PullRequest PullRequest `json:"pull_request"`
}

// Issue entity.
type Issue struct {
	Number int `json:"number"`
}

// PullRequest entity.
type PullRequest struct {
	Number int `json:"number"`
}

// GetEvent from GITHUB_EVENT_PATH.
func GetEvent() (*Event, error) {
	p := os.Getenv("GITHUB_EVENT_PATH")
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	e := &Event{}
	if err := json.NewDecoder(f).Decode(e); err != nil {
		return nil, err
	}
	return e, nil
}

var client *github.Client

// NewClient returns a new github client.
func NewClient(ctx *context.Context, token string) *github.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(*ctx, ts)

	client = github.NewClient(tc)

	return client
}
