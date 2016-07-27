package githubapi

import (
	"errors"
	_ "log"
)

// WebHookEvents ...
type WebHookEvents struct {
	ClosePr         HookHandler
	OpenPr          HookHandler
	EditPr          HookHandler
	AssignedPr      HookHandler
	UnassignedPr    HookHandler
	CloseIssue      HookHandler
	OpenIssue       HookHandler
	EditIssue       HookHandler
	AssignedIssue   HookHandler
	UnassignedIssue HookHandler
}

// HookHandler A function that can do something using the api as a result of a webhook event.
type HookHandler func(WebhookEvent) error

// ErrUnhandledEvent ...
var ErrUnhandledEvent = errors.New("Unhandled event recieved to hook Handlered. Nothing done.")

// HandleWebHookEvents Handle calling function for a webhookEvent...
func (e *WebHookEvents) HandleWebHookEvents(evt WebhookEvent) (err error) {
	switch evt.Type() {
	case prType:
		switch evt.Action {
		case "closed":
			if e.ClosePr != nil {
				return e.ClosePr(evt)
			}
		case "edited":
			if e.EditPr != nil {
				return e.EditPr(evt)
			}
		case "opened":
			if e.OpenPr != nil {
				return e.OpenPr(evt)
			}
		case "assigned":
			if e.AssignedPr != nil {
				return e.AssignedPr(evt)
			}
		case "unassigned":
			if e.UnassignedPr != nil {
				return e.UnassignedPr(evt)
			}
		}
	case issueType:
		switch evt.Action {
		case "closed":
			if e.CloseIssue != nil {
				return e.CloseIssue(evt)
			}
		case "edited":
			if e.EditIssue != nil {
				return e.EditIssue(evt)
			}
		case "opened":
			if e.OpenIssue != nil {
				return e.OpenIssue(evt)
			}
		case "assigned":
			if e.AssignedIssue != nil {
				return e.AssignedIssue(evt)
			}
		case "unassigned":
			if e.UnassignedIssue != nil {
				return e.UnassignedIssue(evt)
			}
		}
	}
	err = ErrUnhandledEvent
	log.Println("Unhandled: ", evt.Type(), evt.Action)
	return
}
