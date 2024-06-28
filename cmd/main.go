package main

import (
	"context"
	"log"
	"os"

	"drake.elearn-platform.ru/internal/configs"
	"drake.elearn-platform.ru/internal/systems"
	"drake.elearn-platform.ru/monoliths/auth"
	onlinecourses "drake.elearn-platform.ru/monoliths/online_courses"
	"github.com/joho/godotenv"
)

type monoliths struct {
	*systems.System
	modules []systems.Module
}

func main() {
	InitReader()
	appConfig := configs.NewAppConfig()
	modules := &monoliths{
		System: systems.NewSystem(appConfig),
		modules: []systems.Module{
			auth.NewAuthModule(),
			onlinecourses.NewOnlineCoursesModule(),
			//more modules, each can be a separate microservice
		},
	}

	modules.StartupModules(modules.modules)

	modules.Waiter().WaitFor(
		modules.WaitForHttpServer,
		modules.WaitForRPC,
	)

	err := modules.Waiter().Wait()
	if err != nil {
		log.Fatalf("Error waiting for all services to start: %v", err)
		os.Exit(1)
	}
}

func (m *monoliths) StartupModules(modules []systems.Module) {
	// Start all modules
	for _, module := range modules {
		err := module.StartUp(context.Background(), m)
		if err != nil {
			log.Fatalf("Error starting module: %v", err)
		}
	}

}

func InitReader() {
	environment := ""
	if len(os.Args) < 2 {
		log.Fatalf("Env not supplied in argument")
	} else {
		environment = os.Args[1]
	}

	err := godotenv.Load(environment + ".env")
	if err != nil {
		log.Fatalf("Error loading %s.env file", environment)
	}
}
