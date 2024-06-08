package books

import (
	"drake.elearn-platform.ru/internal/adapters"
	"drake.elearn-platform.ru/internal/transport/http/routers/requests"
	"drake.elearn-platform.ru/static"
	"github.com/go-chi/chi/v5"
)

type BookRouter interface {
}

type bookRouter struct {
	chi.Router
	method adapters.RequestMethod
	path   string
}

func (g *bookRouter) RequestMethod(method static.HttpMethod) adapters.RequestMethod {
	g.method.Method(method)
	return g.method
}

func (g *bookRouter) Path(path string) adapters.Router {
	g.path = path
	return g
}

func (g *bookRouter) RegisterGroup(path string, r adapters.Router) {
	//TODO implement me
	panic("implement me")
}

func (g *bookRouter) Build() {
	g.RequestMethod(static.GET).
		Register("/:id", g.GetByID()).
		Register("/all", g.ListAll()).
		Register("/filter", g.FindByFilter())
}

func NewBookRouter(r chi.Router) adapters.Router {
	return &bookRouter{Router: r, method: requests.NewRequestMethod()}
}
