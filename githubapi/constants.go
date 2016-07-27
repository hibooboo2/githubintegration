package githubapi

import (
	// Use our custom log
	_ "log"
	"os"
)

var apiToken = os.Getenv("GITHUB_API_TOKEN")

const apiURL = `https://api.github.com`

func init() {
	if apiToken == "" {
		log.Println("Need GITHUB_API_TOKEN defined as an environment variable.")
		os.Exit(1)
	}
}
