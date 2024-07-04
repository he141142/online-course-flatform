package systems

import (
	"context"
	"fmt"
	"net/http"

	"drake.elearn-platform.ru/internal/adapters"
	"drake.elearn-platform.ru/internal/waiter"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // This is the PostgreSQL driver package
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zapadapter"
	"golang.org/x/sync/errgroup"

	"drake.elearn-platform.ru/internal/configs"
	webservers "drake.elearn-platform.ru/internal/web_servers"
	"drake.elearn-platform.ru/pkg/logger"
)

type System struct {
	AppConfig configs.AppConfig
	WebServer *webservers.HttpChiInstance
	DbConn    *sqlx.DB
	waiter    adapters.Waiter
}

var _ Service = (*System)(nil)

func NewSystem(cfg configs.AppConfig) *System {
	system := &System{AppConfig: cfg}
	if cfg.EnablePostgres {
		system.initDatabaseConnection()
	}
	r := chi.NewRouter()
	system.WebServer = webservers.NewChiHttpServer(r)
	system.waiter = waiter.NewWaiter(waiter.CatchSignalCfg(true))
	return system
}

func (sys *System) initDatabaseConnection() {
	dataSource := sys.AppConfig.DBConfig.CreateConnectionString()
	client, err := sqlx.Open(sys.AppConfig.DBConfig.Driver, dataSource)
	if err != nil {
		panic(err)
	}

	logger.InitBasic(
		&logger.LogConfig{
			LogLevel:    sys.AppConfig.DBConfig.LogLevel,
			ServiceName: sys.AppConfig.DBConfig.ServiceName,
		},
	)
	loggerAdapter := zapadapter.New(logger.ZapBasicLogger)
	client = sqlx.NewDb(sqldblogger.OpenDriver(dataSource, client.Driver(),
		loggerAdapter,
		sqldblogger.WithMinimumLevel(sys.AppConfig.DBConfig.MinLevel),
	),
		sys.AppConfig.DBConfig.Driver,
	)

	client.SetMaxIdleConns(sys.AppConfig.DBConfig.MaxIdleConnnection)
	client.SetMaxOpenConns(sys.AppConfig.DBConfig.MaxOpenConnection)
	client.SetConnMaxLifetime(sys.AppConfig.DBConfig.MaxLifeTime)
	client.SetConnMaxIdleTime(sys.AppConfig.DBConfig.MaxLifeTime)
	// verifies connection is db is working
	if err := client.Ping(); err != nil {
		panic(err)
	}
	sys.DbConn = client
}

func (sys *System) WaitForHttpServer(ctx context.Context) error {
	restServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", sys.AppConfig.WebServerConfig.Port),
		Handler: sys.HttpClient(),
	}

	group, grCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		logger.Info("Starting HTTP server", context.Background(), nil)
		return restServer.ListenAndServe()
	})
	group.Go(func() error {
		<-grCtx.Done()
		fmt.Println("web server to be shutdown")
		ctx, cancel := context.WithTimeout(context.Background(), sys.AppConfig.ShutdownTimeout)
		defer cancel()
		if err := restServer.Shutdown(ctx); err != nil {
			return err
		}
		return nil
	})

	return group.Wait()
}

func (sys *System) WaitForRPC(ctx context.Context) error {
	return nil
}

func (sys *System) DbClient() *sqlx.DB {
	return sys.DbConn
}

func (sys *System) HttpClient() *webservers.HttpChiInstance {
	return sys.WebServer
}

func (sys *System) GetAppConfig() configs.AppConfig {
	return sys.AppConfig
}

func (sys *System) Waiter() adapters.Waiter {
	return sys.waiter
}
