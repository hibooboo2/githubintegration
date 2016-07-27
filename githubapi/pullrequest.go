package githubapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// PullRequest ...
type PullRequest struct {
	Body   string           `json:"body"`
	Merged bool             `json:"merged"`
	Base   *PullRequestBase `json:"base"`
	Head   *PullRequestBase `json:"head"`
	Issue
	IssueURL string `json:"issue_url"`
}

// PullRequestBase ...
type PullRequestBase struct {
	Repo  Repo   `json:"repo"`
	Sha   string `json:"sha"`
	Label string `json:"label"`
}

const (
	repoRegex = `\\s([a-zA-Z-]+/[a-zA-Z0-9-]+)?#([0-9]+)`
)

var testRefRegex = regexp.MustCompile("[Tt]est(ing|ed|s)?" + repoRegex)
var refRepoRegex = regexp.MustCompile(repoRegex)

// Get ... Get a pr from the api.
func (pr *PullRequest) Get() (err error) {
	req, err := http.NewRequest("GET", pr.URL, nil)
	if err != nil {
		return
	}
	req.Header.Add("authorization", "token "+apiToken)
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, pr)
	if err != nil {
		log.Println(res.StatusCode)
		log.Println(res)
		log.Println(string(body))
		return
	}
	if !strings.HasSuffix(pr.HTMLURL, fmt.Sprintf("pull/%d", pr.Number)) {
		log.Println("Not PullRequest: ", pr)
		return ErrIsNotPR
	}
	err = pr.Issue.Get()
	return err
}

// SetIssuesToTest ... Called when a pr is closed. Used to label any issues as testing/ to test
func (pr *PullRequest) SetIssuesToTest() (err error) {
	if pr.State != "closed" || !pr.Merged {
		return
	}
	matches := testRefRegex.FindAllStringSubmatch(" "+pr.Body, 200)
	log.Println(pr.Body)
	log.Println(matches)
	issuesTOUpdate := make(map[string]Issue)
	for _, x := range matches {
		if x[2] == "" {
			x[2] = pr.Base.Repo.FullName
		}
		log.Println(x, len(x))
		issue := Issue{URL: fmt.Sprintf(issueURL, x[2], x[3])}
		issuesTOUpdate[issue.URL] = issue
	}
	var errored bool
	for _, issue := range issuesTOUpdate {
		err := issue.SetWaffleStatus(string(testingStatus))
		if !issue.IsPr && err != nil {
			log.Println("Errored when updating label for: ", issue.Number, " Error: ", err.Error())
			errored = true
		}
	}
	if errored {
		err = errors.New("Failed to update one or more issue labels.")
	}
	return
}

// ReferencedIssues ...
func (pr *PullRequest) ReferencedIssues() (issues map[int]Issue, err error) {
	err = pr.Get()
	if err != nil && !pr.IsPr && err != ErrIsNotPR {
		return
	}
	matches := refRepoRegex.FindAllStringSubmatch(" "+pr.Body, 200)
	issues = make(map[int]Issue)
	for _, x := range matches {
		if x[1] == "" {
			x[1] = pr.Base.Repo.FullName
		}
		log.Println(x, len(x))
		issue := Issue{URL: fmt.Sprintf(issueURL, x[1], x[2])}
		issues[issue.Number] = issue
	}
	for _, issue := range issues {
		err = issue.Get()
		if err != nil && !issue.IsPr {
			issues = make(map[int]Issue)
			return
		}
	}
	return
}

// ReferencesAnIssue ...
func (pr *PullRequest) ReferencesAnIssue() (bool, error) {
	issues, err := pr.ReferencedIssues()
	if err != nil {
		// Errored seeing if pr references an Issue.
		return false, err
	}
	return len(issues) > 0, nil
}
