package handlers

import "net/http"

type Api interface {
	Health(w http.ResponseWriter, r *http.Request)
}

type ApiWrapper struct {
	api Api
}

func (aw *ApiWrapper) Health(w http.ResponseWriter, r *http.Request) {
	aw.api.Health(w, r)
}

func SubscribeHandlers(api Api) {
	wrapper := ApiWrapper{
		api,
	}

	http.HandleFunc("/", wrapper.Health)
	http.HandleFunc("/health", wrapper.Health)
}
