package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"drake.elearn-platform.ru/internal/transport/http/routers/modules/books"
	"drake.elearn-platform.ru/static"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
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

	bookRouter := books.NewBookRouter(r)
	//add custom api:
	bookRouter.RequestMethod(static.GET).
		Register("/all", func(writer http.ResponseWriter, request *http.Request) {

		}).
		Register("/all", func(writer http.ResponseWriter, request *http.Request) {

		})

	bookRouter.Build()

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
