package handlers

import (
	"fmt"
	"net/http"
)

type healthHandler struct {
	// (Omitted)
}

// Health returns "OK" if the application is running properly.
func (h healthHandler) Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
