package handlers

type Handler struct {
	healthHandler
	sampleHandler
}

func New() Handler {
	return Handler{}
}
