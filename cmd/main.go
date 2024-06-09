package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

	"drake.elearn-platform.ru/internal/configs"
	"drake.elearn-platform.ru/internal/systems"
)

func main() {
	InitReader()
	r := chi.NewRouter()
	r.Use(cors.Default().Handler)
	if os.Getenv("SECRET") == "" {
		log.Fatalf("SECRET is not defined in the env variable")
	}
	r.Group(func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(map[string]bool{"ok": true})
		})
	})
	// approach := os.Args[2]
	// errgrouppoc.Run(approach)

	appConfig := configs.NewAppConfig()
	_ = systems.NewSystem(appConfig)

	time.Sleep(time.Hour)
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
