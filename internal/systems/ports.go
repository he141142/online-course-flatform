package systems

import (
	"context"

	"drake.elearn-platform.ru/internal/adapters"
	"github.com/jmoiron/sqlx"

	"drake.elearn-platform.ru/internal/configs"
	webservers "drake.elearn-platform.ru/internal/web_servers"
)

type Service interface {
	DbClient() *sqlx.DB
	HttpClient() *webservers.HttpChiInstance
	GetAppConfig() configs.AppConfig
	Waiter() adapters.Waiter
}

type Module interface {
	StartUp(ctx context.Context, service Service) error
}
