package genres

import (
	"drake.elearn-platform.ru/internal/adapters"
	"drake.elearn-platform.ru/internal/transport/http/routers/requests"
	"drake.elearn-platform.ru/static"
	"github.com/go-chi/chi/v5"
)

type genresRouter struct {
	chi.Router
	method adapters.RequestMethod
	path   string
}

func (g *genresRouter) RequestMethod(
	method static.HttpMethod) adapters.RequestMethod {
	g.method.Method(method)
	return g.method
}

func (g *genresRouter) Path(path string) adapters.Router {
	g.path = path
	return g
}

func (g *genresRouter) RegisterGroup(path string, r adapters.Router) {
	//TODO implement me
	panic("implement me")
}

func (g *genresRouter) Build() {
	g.RequestMethod(static.GET).
		Register("/:id", g.GetGenresByID()).
		Register("/all", g.ListAll()).
		Register("/filter", g.FindByFilter())

}

func (g *genresRouter) apply() {
	g.Build()
	g.RegisterGroup(g.path, g)
	g.Get("/", g.ListAll())
}

func NewGenresRouter(
	r chi.Router) adapters.Router {
	return &genresRouter{Router: r, method: requests.NewRequestMethod()}
}
