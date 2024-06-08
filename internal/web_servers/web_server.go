package webservers

import (
	"fmt"

	"drake.elearn-platform.ru/static"
)

type HttpServer struct {
	Protocol           string
	Host               string
	Port               int
	ApiVersion         string
	WebServerFramework static.WebServerFramework
}

func (http HtHttpServer) GetURL() string {
	return fmt.Sprintf("%s")
}