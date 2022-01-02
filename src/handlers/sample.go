package handlers

import (
	"fmt"
	"net/http"
	"time"
)

type sampleHandler struct {
	// (Omitted)
}

// Heavy performs a busy loop.
func (h sampleHandler) Heavy(w http.ResponseWriter, r *http.Request) {
	busyloop()
	fmt.Fprintf(w, "SUCCESS: heavy")
}

// Sleep performs a sleep for 10 min.
func (h sampleHandler) Sleep(w http.ResponseWriter, r *http.Request) {
	sleep()
	fmt.Fprintf(w, "SUCCESS: sleep")
}

func busyloop() {
	for i := 0; i < 1000; i++ {
		foo(1000)
	}
}

func foo(scale int) {
	baz(scale)
	bar(scale)
}

func baz(scale int) {
	load(scale)
}

func bar(scale int) {
	load(scale * 2)
}

func load(scale int) {
	for i := 0; i < scale*(1<<14); i++ {
	}
}

func sleep() {
	time.Sleep(600 * time.Second)
}
