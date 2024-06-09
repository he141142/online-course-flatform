package systems

import (
	"github.com/jmoiron/sqlx"

	"drake.elearn-platform.ru/internal/configs"
	webservers "drake.elearn-platform.ru/internal/web_servers"
)

type System struct {
	AppConfig configs.AppConfig
	WebServer *webservers.HttpChiInstance
	DbConn    *sqlx.DB
}

func NewSystem(cfg configs.AppConfig) *System {
	return &System{AppConfig: cfg}
}