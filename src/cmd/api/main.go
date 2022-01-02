package main

import (
	"log"
	"net/http"
	"os"

	"github.com/hyorimitsu/hello-cloud-profiler/src/gcp"
	"github.com/hyorimitsu/hello-cloud-profiler/src/gcp/profiler"
	"github.com/hyorimitsu/hello-cloud-profiler/src/handlers"
)

func main() {
	code := start()
	os.Exit(code)
}

func start() int {
	if gcp.RunningOnGCP() {
		p := profiler.NewDefault()
		if err := p.Start(); err != nil {
			log.Println(err)
		}
	}

	handler := handlers.New()
	handlers.SubscribeHandlers(handler)

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		return 1
	}
	return 0
}
