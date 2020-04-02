package cmd

import (
	"context"
	"fmt"

	"github.com/jonyhy96/lgtm-action/pkg/runner"
	"github.com/jonyhy96/lgtm-action/pkg/util/github"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// Execute command
func Execute(ctx *context.Context, version string) {
	input := new(Input)
	var (
		rootCmd = &cobra.Command{
			Use:     "lgtm action command",
			Short:   "lgtm github action helps you to approve",
			Version: version,
			RunE:    newRunCommand(ctx, input),
		}
	)
	rootCmd.Flags().StringVarP(&input.GithubAuthToken, "token", "g", "", "the GITHUB_AUTH_TOKEN to make changes on your repo")
	rootCmd.Flags().IntVarP(&input.Times, "times", "t", 1, "the port of the server (default 8080)")
	rootCmd.Flags().StringVarP(&input.OwnersFile, "owners", "o", "OWNERS", "the owners file that contains the repo's owner")
	if err := rootCmd.Execute(); err != nil {
		logrus.Error(fmt.Errorf("rootCmd Execute error: %s", err.Error()))
	}
}

func newRunCommand(ctx *context.Context, input *Input) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		var (
			client = github.NewClient(ctx, input.GithubAuthToken)
			runner = runner.New(ctx, client)
		)
		return runner.Run(input.Times, input.OwnersFile)
	}
}
