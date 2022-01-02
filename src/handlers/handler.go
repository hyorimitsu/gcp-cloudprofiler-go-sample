package handlers

// Handler is api handlers.
type Handler struct {
	healthHandler
	sampleHandler
}

// New returns new Handler.
func New() Handler {
	h := Handler{}

	// (Omitted)

	return h
}
