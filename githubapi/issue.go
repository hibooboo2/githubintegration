package githubapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	// Use custom log
	_ "log"
	"net/http"
	"strings"
)

// Issue Github Issue.
type Issue struct {
	URL     string  `json:"url"`
	Labels  []Label `json:"labels"`
	Number  int     `json:"number"`
	HTMLURL string  `json:"html_url"`
	State   string  `json:"state"`
	IsPr    bool    `json:"-"`
}

const issueURL = apiURL + `/repos/%s/issues/%s`

// ErrIsNotIssue ...
var ErrIsNotIssue = errors.New("This is not an Issue")

// ErrIsNotPR ...
var ErrIsNotPR = errors.New("This is not a PR")

// Get ...
func (i *Issue) Get() (err error) {
	log.Println("Start Number:", i.Number, " HTML:", i.HTMLURL, " URL:", i.URL)
	req, err := http.NewRequest("GET", i.URL, nil)
	if err != nil {
		log.Debug(err)
		return
	}

	req.Header.Add("authorization", "token "+apiToken)
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Debug(err)
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Debug(err)
		return
	}
	err = json.Unmarshal(body, i)
	if err != nil {
		log.Debug(err)
		log.Debugln(res.StatusCode)
		log.Debugln(res)
		log.Debugln(string(body))
		return
	}
	if !strings.HasSuffix(i.HTMLURL, fmt.Sprintf("issues/%d", i.Number)) {
		if !strings.HasSuffix(i.HTMLURL, fmt.Sprintf("pull/%d", i.Number)) {
			log.Debugln("Is not issue or pr: ", i.HTMLURL, " Number: ", i.Number)
			return ErrIsNotIssue
		}
		i.IsPr = true
	}
	return
}

// AddLabel ...
func (i *Issue) AddLabel(add string) error {
	var labels struct {
		Labels []string `json:"labels"`
	}
	err := i.Get()
	if err != nil {
		log.Debug(err)
		return err
	}
	labels.Labels = []string{add}
	for _, label := range i.Labels {
		labels.Labels = append(labels.Labels, label.Name)
	}
	payLd, err := json.Marshal(&labels)
	if err != nil {
		log.Debug(err)
		return err
	}
	payload := strings.NewReader(string(payLd))

	req, _ := http.NewRequest("PATCH", i.URL, payload)

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

// waffleStatus string
type waffleStatus string

// IsWaffleStatus ...
func IsWaffleStatus(s string) bool {
	status := waffleStatus(s)
	switch status {
	case readyStatus:
		return true
	case inProgressStatus:
		return true
	case testingStatus:
		return true
	default:
		return false
	}
}

const (
	readyStatus      waffleStatus = `ready`
	inProgressStatus waffleStatus = `in progress`
	testingStatus    waffleStatus = `testing`
)

// SetWaffleStatus ...
func (i *Issue) SetWaffleStatus(status string) error {
	if !IsWaffleStatus(status) {
		return errors.New("Cannot set status not valid waffle status.")
	}
	var labels struct {
		Labels []string `json:"labels"`
	}
	err := i.Get()
	if err != nil {
		log.Debugln(err)
		return err
	}
	labels.Labels = []string{string(status)}
	log.Debug("Labels before: ")
	for _, label := range i.Labels {
		log.Debug("$     label: '%s'", label.Name)
	}
	for _, label := range i.Labels {
		if IsWaffleStatus(label.Name) {
			log.Debug("Dropping ", label.Name)
			break
		}
		log.Debug("Keeping ", label.Name)
		labels.Labels = append(labels.Labels, label.Name)
	}
	log.Debug("Labels after: ")
	for _, label := range labels.Labels {
		log.Debug("$     label: '%s'", label)
	}
	payLd, err := json.Marshal(&labels)
	if err != nil {
		return err
	}
	payload := strings.NewReader(string(payLd))

	req, _ := http.NewRequest("PATCH", i.URL, payload)

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
