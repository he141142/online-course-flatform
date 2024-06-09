package webservers

import (
	"fmt"

	"drake.elearn-platform.ru/static"
)

type HttpServerConfig struct {
	Protocol           string
	Host               string
	Port               int
	ApiVersion         string
	WebServerFramework static.WebServerFramework
}

func (http HttpServerConfig) GetURL() string {
	return fmt.Sprintf("%s://%s:%d", http.Protocol, http.Host, http.Port)
}

type WebServerInstance interface {
}