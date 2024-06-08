package configs

import (
	"time"

	webservers "drake.elearn-platform.ru/internal/web_servers"
)

type AppConfig struct {
	ShutdownTimeout time.Duration
	AppName         string
	Appversion      string
	WebServerConfig webservers.HttpServerConfig
}