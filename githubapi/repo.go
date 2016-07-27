package githubapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	// Use our custom log.
	_ "log"
	"net/http"
	"strings"
)

// Repo ...
type Repo struct {
	FullName    string `json:"full_name"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
	Fork        bool   `json:"fork"`
	URL         string `json:"url"`
	HTMLURL     string `json:"html_url"`
}

// Get ...
func (r *Repo) Get() (err error) {
	req, err := http.NewRequest("GET", r.URL, nil)
	if err != nil {
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
	err = json.Unmarshal(body, r)
	if err != nil {
		log.Debug(err)
		log.Debugln(res.StatusCode)
		log.Debugln(res)
		log.Debugln(string(body))
		return
	}
	if !strings.HasSuffix(r.HTMLURL, fmt.Sprintf("repos/%s", r.FullName)) {
		err = errors.New("Not a Repo: " + r.HTMLURL)
		log.Debug(err)
		return
	}
	return
}
