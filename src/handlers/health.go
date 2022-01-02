package handlers

import (
	"fmt"
	"net/http"
)

type healthHandler struct {
	// (Omitted)
}

func (h healthHandler) Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
