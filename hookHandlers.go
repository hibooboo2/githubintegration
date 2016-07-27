package main

import (
	"log"
	"net/http"
	"time"

	"github.com/codehack/go-relax"
	api "github.com/hibooboo2/githubissues/githubapi"
)

// WebHookServer ...
type WebHookServer struct {
	Hooks *api.WebHookEvents
}

// Index ...
func (e *WebHookServer) Index(ctx *relax.Context) {
	log.Println("Index")
	ctx.Respond(nil, http.StatusOK)
}

// HandleWebHookServer ...
func (e *WebHookServer) HandleWebHookServer(ctx *relax.Context) {
	evt := api.WebhookEvent{}
	err := ctx.Decode(ctx.Request.Body, &evt)
	if err != nil {
		ctx.Respond(err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Event: ", evt.Type(), ".", evt.Action)
	if e.Hooks == nil {
		ctx.Respond(nil, http.StatusInternalServerError)
	}
	err = e.Hooks.HandleWebHookEvents(evt)
	if err != nil {
		log.Println(err.Error())
		unhandled(ctx, evt)
		return
	}
	ctx.Respond(nil, http.StatusOK)
}

func unhandled(ctx *relax.Context, evt api.WebhookEvent) {
	log.Println("Unsupported Event: ", evt.Type(), ".", evt.Action)
	// data, _ := json.MarshalIndent(evt, "    ", "")
	// log.Println(string(data))
	ctx.Respond(nil, http.StatusOK)
}

func statusCheck(evt api.WebhookEvent) error {
	err := evt.PullRequest.Get()
	if err != nil {
		log.Println("Open Pr GetPR: ", err.Error())
		return err
	}
	c := api.Commit{Sha: evt.PullRequest.Head.Sha, RepoFullName: evt.PullRequest.Base.Repo.FullName}
	go func() {
		time.Sleep(time.Second * 10)
		c.CreateStatus("success", "http://github.jhrb.us", "github/jhrb/integration", "Pr just opened")
	}()
	return c.CreateStatus("pending", "http://github.jhrb.us", "github/jhrb/integration", "Pr just opened")
}
