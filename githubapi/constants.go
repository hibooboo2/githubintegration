package githubapi

import "os"

var apiToken = os.Getenv("GITHUB_API_TOKEN")

const apiURL = `https://api.github.com`

func init() {
	if apiToken == "" {
		panic("Need GITHUB_API_TOKEN defined as an environment variable.")
	}
}
