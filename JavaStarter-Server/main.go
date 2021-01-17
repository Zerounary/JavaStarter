package main

import (
	"log"
	"net/http"

	"github.com/sosedoff/gitkit"
)

func main() {
	// Configure git hooks
	// hooks := &gitkit.HookScripts{
	// 	PreReceive: `echo "Hello World!"`,
	// }

	// Configure git service
	service := gitkit.New(gitkit.Config{
		Dir:        "/Users/mac/go-git-http",
		AutoCreate: true,
		AutoHooks:  true,
		// Hooks:      hooks,
	})

	// Configure git server. Will create git repos path if it does not exist.
	// If hooks are set, it will also update all repos with new version of hook scripts.
	if err := service.Setup(); err != nil {
		log.Fatal(err)
	}

	http.Handle("/", service)

	// Start HTTP server
	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal(err)
	}
}
