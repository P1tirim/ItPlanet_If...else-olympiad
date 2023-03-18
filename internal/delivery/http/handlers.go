package http

import "api/internal/services"

type Handlers struct {
	services *services.Services
}

func MakeHandlers(apiservices *services.Services) *Handlers {
	return &Handlers{apiservices}
}
