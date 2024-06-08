package requests

import (
	"net/http"

	"drake.elearn-platform.ru/internal/adapters"
	"drake.elearn-platform.ru/static"
)

type requestMethod struct {
	curr      static.HttpMethod
	methodMap map[static.HttpMethod]func(path string, hdl http.HandlerFunc)
}

func (c *requestMethod) Method(method static.HttpMethod) adapters.RequestMethod {
	c.curr = method
	return c
}

func (c *requestMethod) Register(path string, hdl http.HandlerFunc) adapters.RequestMethod {
	c.methodMap[c.curr](path, hdl)
	return c
}

func (c *requestMethod) validate() {
	if c.methodMap == nil {
		c.methodMap = make(map[static.HttpMethod]func(path string, hdl http.HandlerFunc))
	}
}
func NewRequestMethod() adapters.RequestMethod {
	return &requestMethod{
		curr:      static.GET,
		methodMap: make(map[static.HttpMethod]func(path string, hdl http.HandlerFunc)),
	}
}
