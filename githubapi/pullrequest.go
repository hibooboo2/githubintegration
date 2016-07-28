package githubapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	// Use our custom log
	_ "log"
	"net/http"
	"regexp"
	"strconv"
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

var testRefRegex = regexp.MustCompile("[Tt]est(ing|ed|s)?\\s([a-zA-Z-]+/[a-zA-Z0-9-]+)?#([0-9]+)")
var refRepoRegex = regexp.MustCompile("\\s([a-zA-Z-]+/[a-zA-Z0-9-]+)?#([0-9]+)")

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
	pr.Get()
	if pr.State != "closed" && pr.Merged {
		return
	}
	matches := testRefRegex.FindAllStringSubmatch(" "+pr.Body, 200)
	log.Debug(pr.Body)
	log.Debugln(matches)
	issuesTOUpdate := make(map[string]Issue)
	for _, x := range matches {
		log.Debugln("Found : ", len(x), " In: ", x)
		if x[2] == "" {
			x[2] = pr.Base.Repo.FullName
		}
		log.Debug(x, len(x))
		issue := Issue{URL: fmt.Sprintf(issueURL, x[2], x[3])}
		issuesTOUpdate[issue.URL] = issue
	}
	var errored bool
	for _, issue := range issuesTOUpdate {
		err := issue.SetWaffleStatus(string(testingStatus))
		if !issue.IsPr && err != nil {
			log.Debugln("Errored when updating label for: ", issue.Number, " Error: ", err.Error())
			errored = true
		}
	}
	if errored {
		err = errors.New("Failed to update one or more issue labels.")
		log.Debug(err)
	}
	return
}

// ReferencedIssues ...
func (pr *PullRequest) ReferencedIssues() (issues map[int]Issue, err error) {
	issues = make(map[int]Issue)
	log.Debugln("Pr is:", pr)
	err = pr.Get()
	if err != nil && !pr.IsPr && err != ErrIsNotPR {
		log.Debug(err)
		return
	}
	matches := refRepoRegex.FindAllStringSubmatch(" "+pr.Body, 200)
	log.Debugln("Body are: ", pr.Body)
	log.Debugln("Matches are: ", matches)
	if len(matches) == 0 {
		err = errors.New("No issue refs found")
		log.Debugln(err)
		return
	}
	for _, x := range matches {
		if x[1] == "" {
			x[1] = pr.Base.Repo.FullName
		}
		log.Debugln("x from matches: ", x, len(x))
		issue := Issue{URL: fmt.Sprintf(issueURL, x[1], x[2])}
		issue.Number, err = strconv.Atoi(x[2])
		if _, there := issues[issue.Number]; there {
			break
		}
		if err != nil {
			log.Debugln(err)
			issues = make(map[int]Issue)
			return
		}
		err = issue.Get()
		if err != nil && !issue.IsPr {
			issues = make(map[int]Issue)
			log.Debug(err)
			return
		}
		if !issue.IsPr {
			issues[issue.Number] = issue
		}
	}
	return
}

// ReferencesAnIssue ...
func (pr *PullRequest) ReferencesAnIssue() (bool, error) {
	log.Debugln("Called ReferencesAnIssue")
	issues, err := pr.ReferencedIssues()
	if err != nil {
		log.Debug(err)
		// Errored seeing if pr references an Issue.
		return false, err
	}
	return len(issues) > 0, nil
}
