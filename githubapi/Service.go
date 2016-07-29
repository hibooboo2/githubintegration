package githubapi

// Service Defines a service that uses the api to do things.
type Service interface {
	Add(o *OauthToken) error
	Delete(o *OauthToken) error
	Pause(o *OauthToken) error
	Event(evt WebhookEvent) error
	Name() error
}
