package env

import (
	"os"
	"strings"
)

// GITHUB_ prefix env.
var (
	GithubRef        = os.Getenv("GITHUB_REF")
	GithubEventName  = os.Getenv("GITHUB_EVENT_NAME")
	GithubRepository = os.Getenv("GITHUB_REPOSITORY")
	Owner            = strings.Split(GithubRepository, "/")[0]
	Repository       = strings.Split(GithubRepository, "/")[1]
)
