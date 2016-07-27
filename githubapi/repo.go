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
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, r)
	if err != nil {
		log.Println(res.StatusCode)
		log.Println(res)
		log.Println(string(body))
		return
	}
	if !strings.HasSuffix(r.HTMLURL, fmt.Sprintf("repos/%s", r.FullName)) {
		err = errors.New("Not a Repo: " + r.HTMLURL)
		return
	}
	return
}
