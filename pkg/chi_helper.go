package pkg

import (
	"net/http"

	"drake.elearn-platform.ru/static"
	"github.com/go-chi/chi/v5"
)

func ByMethod(route chi.Router) map[static.HttpMethod]func(path string, hdl http.HandlerFunc) {
	return map[static.HttpMethod]func(path string, hdl http.HandlerFunc){
		static.GET:    route.Get,
		static.POST:   route.Post,
		static.PUT:    route.Put,
		static.DELETE: route.Delete,
	}
}
