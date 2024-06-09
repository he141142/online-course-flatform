package configs

import (
	"os"
	"time"

	webservers "drake.elearn-platform.ru/internal/web_servers"
	"drake.elearn-platform.ru/pkg/utils"
)

type AppConfig struct {
	ShutdownTimeout time.Duration
	AppName         string
	Appversion      string
	WebServerConfig webservers.HttpServerConfig
	DBConfig        DBConfig
}

func (c *AppConfig) initDbConfig() {
	hostName := os.Getenv("DB_HOST")
	port := utils.EnvGetIntValue("DB_PORT", 7888)
	userName := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSchema := os.Getenv("DB_SCHEMA")
	driver := os.Getenv("DB_DRIVER")
	logLevel := os.Getenv("LOG_LEVEL")
	serviceName := os.Getenv("SERVICE_NAME")
	maxIdleConnection := utils.EnvGetIntValue("DB_MAX_IDLE_CONNECTION", 10)
	maxLifeTime := utils.EnvGetIntValue("DB_CONN_MAX_LIFETIME_SECS", 600)
	maxOpenConnection := utils.EnvGetIntValue("DB_MAX_OPEN_CONNECTION", 10)
	dbConfig := DBConfig{
		Host:               hostName,
		Port:               port,
		User:               userName,
		Password:           password,
		DBName:             dbName,
		DbSchema:           dbSchema,
		Driver:             driver,
		LogLevel:           logLevel,
		ServiceName:        serviceName,
		MaxIdleConnnection: maxIdleConnection,
		MaxLifeTime:        time.Duration(maxLifeTime) * time.Second,
		MaxOpenConnection:  maxOpenConnection,
	}
	c.DBConfig = dbConfig
}

func NewAppConfig() AppConfig {
	appName := os.Getenv("APP_NAME")
	appVersion := os.Getenv("APP_VERSION")
	shutdownTimeout := utils.EnvGetIntValue("APP_SHUTTDOWN_TIMEOUT_IN_SECS", 30)

	appCfg := AppConfig{
		AppName:         appName,
		Appversion:      appVersion,
		ShutdownTimeout: time.Duration(shutdownTimeout) * time.Second,
	}

	appCfg.initDbConfig()

	return appCfg
}