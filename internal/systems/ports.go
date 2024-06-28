package systems

import (
	"context"

	"github.com/jmoiron/sqlx"

	"drake.elearn-platform.ru/internal/configs"
	webservers "drake.elearn-platform.ru/internal/web_servers"
)

type Service interface {
	DbClient() *sqlx.DB
	HttpClient() *webservers.HttpChiInstance
	GetAppConfig() configs.AppConfig
}

type Module interface {
	StartUp(ctx context.Context, service Service) error
}