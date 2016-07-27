package githubapi

// WebhookEvent ...
type WebhookEvent struct {
	Action      string       `json:"action"`
	Issue       *Issue       `json:"issue,omitempty"`
	PullRequest *PullRequest `json:"pull_request,omitempty"`
	Repo        *Repo        `json:"repository"`
}

const (
	// PRType denotes that an event is a pr.
	prType = "pullrequest"
	// IssueType denotes that an event is an issue.
	issueType = "issue"
)

// Type Get the type of event.
func (evt *WebhookEvent) Type() string {
	switch {
	case evt.PullRequest != nil:
		return prType
	case evt.Issue != nil:
		return issueType
	default:
		return "unknown"
	}
}
