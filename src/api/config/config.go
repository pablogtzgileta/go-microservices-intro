package config

import (
	"fmt"
	"os"
)

const (
	apiGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
)

var (
	githubAccessToken = os.Getenv(apiGithubAccessToken)
)

func GetGithubAccessToken() string {
	fmt.Println("token", githubAccessToken)
	return githubAccessToken
}
