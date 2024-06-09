package configs

import (
	"fmt"
	"time"

	sqldblogger "github.com/simukti/sqldb-logger"
)

type DBConfig struct {
	Host       string
	Port       int
	SSLDisable bool
	User       string
	Password   string
	DBName     string
	DbSchema   string
	Driver     string

	LogLevel    string
	ServiceName string

	MaxIdleConnnection int
	MaxLifeTime        time.Duration
	MaxOpenConnection  int
	MinLevel           sqldblogger.Level
}

func (dbCfg DBConfig) CreateConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable binary_parameters=yes",
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.User,
		dbCfg.Password,
		dbCfg.DBName,
	)
}
