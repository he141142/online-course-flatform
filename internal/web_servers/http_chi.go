package webservers

import "github.com/go-chi/chi/v5"

type HttpChiInstance struct {
	chi.Router
}

func NewChiHttpServer(chiRoute chi.Router) WebServerInstance {
	return &HttpChiInstance{Router: chiRoute}
}