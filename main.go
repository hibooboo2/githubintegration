package main

import (
	"log"
	"time"

	api "github.com/hibooboo2/githubissues/githubapi"

	"github.com/codehack/go-relax"
)

func main() {
	log.SetFlags(log.Lshortfile)
	webHookServer()
}

func webHookServer() {
	svc := relax.NewService("/")
	svc.Use(&LogFilter{})

	{
		this := &WebHookServer{}
		hh := &api.WebHookEvents{}
		this.Hooks = hh
		hh.ClosePr = func(evt api.WebhookEvent) error {
			err := evt.PullRequest.SetIssuesToTest()
			if err != nil {
				return err
			}
			return nil
		}
		hh.OpenPr = statusCheck
		hh.EditPr = statusCheck

		r := svc.Resource(this)
		r.POST("/", this.HandleWebHookServer)
	}

	svc.Run(":5554")
}

// LogFilter ...
type LogFilter struct{}

// Run ...
func (l *LogFilter) Run(next func(*relax.Context)) func(*relax.Context) {
	return func(ctx *relax.Context) {
		start := time.Now()
		next(ctx)
		log.Printf("Took: %v\n %s: %s", time.Since(start), ctx.Request.Method, ctx.Request.URL.String())
	}
}
