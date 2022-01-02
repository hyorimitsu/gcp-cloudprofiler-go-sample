package handlers

import (
	"fmt"
	"net/http"
	"time"
)

type sampleHandler struct {
	// (Omitted)
}

func (h sampleHandler) Heavy(w http.ResponseWriter, r *http.Request) {
	// busy loop
	scale := 1000
	for i := 0; i < scale*(1<<14); i++ {
		for j := 0; j < scale*(1<<14); j++ {
			for k := 0; k < scale*(1<<14); k++ {
			}
		}
	}
	fmt.Fprintf(w, "SUCCESS: heavy")
}

func (h sampleHandler) Sleep(w http.ResponseWriter, r *http.Request) {
	time.Sleep(600 * time.Second)
	fmt.Fprintf(w, "SUCCESS: sleep")
}
