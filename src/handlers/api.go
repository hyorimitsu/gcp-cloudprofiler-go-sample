package handlers

import "net/http"

// Api defined all api handlers.
type Api interface {
	Health(w http.ResponseWriter, r *http.Request)
	Heavy(w http.ResponseWriter, r *http.Request)
	Sleep(w http.ResponseWriter, r *http.Request)
}

// ApiWrapper is an Api wrapper to do some processing before and after the actual processing of api.
type ApiWrapper struct {
	api Api
}

// Health calls handler.
// If some processing is to be done before or after the actual processing, it should be described here.
func (aw *ApiWrapper) Health(w http.ResponseWriter, r *http.Request) {
	aw.api.Health(w, r)
}

// Heavy calls handler.
// If some processing is to be done before or after the actual processing, it should be described here.
func (aw *ApiWrapper) Heavy(w http.ResponseWriter, r *http.Request) {
	aw.api.Heavy(w, r)
}

// Sleep calls handler.
// If some processing is to be done before or after the actual processing, it should be described here.
func (aw *ApiWrapper) Sleep(w http.ResponseWriter, r *http.Request) {
	aw.api.Sleep(w, r)
}

// SubscribeHandlers subscribe all handlers
func SubscribeHandlers(api Api) {
	wrapper := ApiWrapper{
		api,
	}

	http.HandleFunc("/", wrapper.Health)
	http.HandleFunc("/health", wrapper.Health)
	http.HandleFunc("/heavy", wrapper.Heavy)
	http.HandleFunc("/sleep", wrapper.Sleep)
}
