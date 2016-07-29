package githubapi

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"html"
	"io/ioutil"
	"net/http"
	"os"
)

const oauthURL = `https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s&state=%s&scope=%s`

var tokens = make(map[string]*OauthToken)

var tokenChannel = make(chan *OauthToken)

var services = []Service{}

var clientID = os.Getenv("OAUTH_GITHUB_CLIENT_ID")
var clientSecret = os.Getenv("OAUTH_GITHUB_SECRET")

// OauthRedirectEndPoint for the application to use when doing oauth redirects.
const OauthRedirectEndPoint = `/oauth`

// AppHost Host that redirect is on/ Servcices probably.
const AppHost = `https://knife.wizardofmath.host`

// OauthRedirectURL full redirect url for oauth.
const OauthRedirectURL = AppHost + OauthRedirectEndPoint

func init() {
	if clientID == "" || clientSecret == "" {
		println("Must provide OAUTH_GITHUB_CLIENT_ID and OAUTH_GITHUB_SECRET as environment variables")
		os.Exit(1)
	}
	http.HandleFunc(OauthRedirectEndPoint, oauthHandler)
	http.HandleFunc("/oauth/add", addUser)
	go func() {
		oauthToken := <-tokenChannel
		log.Infof("New Oauth Token: %##v", *oauthToken)
		for _, svc := range services {
			svc.Add(oauthToken)
		}
	}()
}

// OauthToken Used to authenticate to github for a user.
// There is an automatic callback handler for oauth. When the user has an oauth flow done the token is populated in the toke that was created.
type OauthToken struct {
	token string
	User  User
	state string
	Scopes
}

// AuthorizeURL Get the url used to send user to oauth flow to auth the application you are building.
func (o *OauthToken) AuthorizeURL() string {
	o.state = uuid.NewV4().String()
	tokens[o.state] = o
	if o.Scopes.storedScopes == nil && len(o.Scopes.storedScopes) < 1 {
		o.Scopes.storedScopes = make(map[string]bool)
		o.Scopes.RepoStatus().AdminRepoHook().UserEmail().Repo().AdminOrgHook().AdminRepoHook()
	}
	log.Debugln(o.Scopes.String())
	return html.EscapeString(fmt.Sprintf(oauthURL, clientID, OauthRedirectURL, o.state, o.Scopes.String()))
}

func (o *OauthToken) getAccessToken(state, code string) error {
	return nil
}

func oauthHandler(w http.ResponseWriter, req *http.Request) {
	state := req.URL.Query().Get("state")
	code := req.URL.Query().Get("code")
	oauthToken, ok := tokens[state]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid state provided. Probably being man in the middled."))
		return
	}

	err := oauthToken.getOauthTokenFromGithub(code)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`'{"error":"failed to get Oauth token from github."}'`))
	}
	w.WriteHeader(http.StatusCreated)
	tokenChannel <- oauthToken
}

func addUser(w http.ResponseWriter, req *http.Request) {
	token := OauthToken{}
	token.Scopes.storedScopes = make(map[string]bool)
	w.Write([]byte(fmt.Sprintf(`
        <html>
            <body>
                <a href="%s"><button>Add this integration<button></a>
            </body>
        </html>
        `, token.AuthorizeURL())))
}

func (o *OauthToken) getOauthTokenFromGithub(code string) error {
	url := fmt.Sprintf(`https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s&state=%s`, clientID, clientSecret, code, o.state)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var payload struct {
		AccessToken string `json:"access_token"`
		Scope       string `json:"scope"`
		TokenType   string `json:"token_type"`
	}
	err = json.Unmarshal(body, &payload)
	if err != nil {
		return err
	}

	o.token = payload.AccessToken
	o.Scopes.fromString(payload.Scope)

	return nil
}
