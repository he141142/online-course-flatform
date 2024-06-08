package adapters

import (
	"context"
	"net/http"

	"drake.elearn-platform.ru/static"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type HttpServer interface {
	ListenAndServe(ctx context.Context) error
}

type Router interface {
	Path(path string) Router
	RegisterGroup(path string, r Router)
	RequestMethod(method static.HttpMethod) RequestMethod
	Build()
}

type RequestMethod interface {
	Method(method static.HttpMethod) RequestMethod
	Register(path string, hdl http.HandlerFunc) RequestMethod
}
