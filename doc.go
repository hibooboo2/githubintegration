package main

import (
	"log"
	"os"
)

// Build of githubissues.
var Build = `HEAD`

func init() {
	for _, x := range os.Args {
		if x == "-v" || x == "-version" || x == "--version" {
			log.Println("Build: ", Build)
			os.Exit(0)
		}
	}
}
