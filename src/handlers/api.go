package handlers

type Api interface {
}

type ApiWrapper struct {
	api Api
}

func SubscribeHandlers(api Api) {
	_ = ApiWrapper{
		api,
	}
}
