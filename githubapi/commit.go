package githubapi

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Commit ...
type Commit struct {
	URL          string `json:"url"`
	Sha          string `json:"sha"`
	RepoFullName string `json:"-"`
	HTMLURL      string `json:"html_url"`
}

const (
	statusURL = apiURL + `/repos/%s/statuses/%s`
)

// CreateStatus Set A failed status on commit sha.
// Status must be pending, success, error, failure
func (c *Commit) CreateStatus(status, targetURL, context, description string) error {
	switch status {
	case "pending":
	case "success":
	case "error":
	case "failure":
	default:
		log.Debug("Invalid status for a commit status create passed: ", status)
		return errors.New("Invalid Status provided")
	}
	payload := strings.NewReader(fmt.Sprintf(`
		{
			"state": "%s",
			"target_url": "%s",
			"description": "%s",
			"context": "%s"
		}`, status, targetURL, description, context))

	req, err := http.NewRequest("POST", fmt.Sprintf(statusURL, c.RepoFullName, c.Sha), payload)
	if err != nil {
		log.Debug(err)
		return err
	}

	req.Header.Add("authorization", "token "+apiToken)
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Debug(err)
		return err
	}
	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Debug(err)
		return err
	}
	return nil
}
