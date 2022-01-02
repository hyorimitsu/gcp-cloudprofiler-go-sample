package handlers

import "net/http"

type Api interface {
	Health(w http.ResponseWriter, r *http.Request)
	Heavy(w http.ResponseWriter, r *http.Request)
	Sleep(w http.ResponseWriter, r *http.Request)
}

type ApiWrapper struct {
	api Api
}

func (aw *ApiWrapper) Health(w http.ResponseWriter, r *http.Request) {
	aw.api.Health(w, r)
}

func (aw *ApiWrapper) Heavy(w http.ResponseWriter, r *http.Request) {
	aw.api.Heavy(w, r)
}

func (aw *ApiWrapper) Sleep(w http.ResponseWriter, r *http.Request) {
	aw.api.Sleep(w, r)
}

func SubscribeHandlers(api Api) {
	wrapper := ApiWrapper{
		api,
	}

	http.HandleFunc("/", wrapper.Health)
	http.HandleFunc("/health", wrapper.Health)
	http.HandleFunc("/heavy", wrapper.Heavy)
	http.HandleFunc("/sleep", wrapper.Sleep)
}
