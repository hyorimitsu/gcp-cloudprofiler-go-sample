package handlers

type Handler struct {
	healthHandler
}

func New() Handler {
	return Handler{}
}
