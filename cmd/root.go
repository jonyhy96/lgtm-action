package cmd

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/jonyhy96/lgtm-action/pkg/runner"
	"github.com/jonyhy96/lgtm-action/pkg/util/github"
	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// Execute command
func Execute(ctx *context.Context, version string) {
	input := &Input{
		Times:  "1",      // default times.
		Owners: "OWNERS", // defalt owners.
	}
	var (
		rootCmd = &cobra.Command{
			Use:     "lgtm action command",
			Short:   "lgtm github action helps you to approve",
			Version: version,
			RunE:    newRunCommand(ctx, input),
		}
	)
	input.LoadFromEnv()
	if err := rootCmd.Execute(); err != nil {
		logrus.Error(fmt.Errorf("rootCmd Execute error: %s", err.Error()))
		os.Exit(1)
	}
}

func newRunCommand(ctx *context.Context, input *Input) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		var (
			client = github.NewClient(ctx, input.GithubAuthToken)
			runner = runner.New(ctx, client)
		)
		times, err := strconv.Atoi(input.Times)
		if err != nil {
			return err
		}
		return runner.Run(times, input.Owners)
	}
}
