package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	// load envs from .env file
	godotenv.Load(".env")
	portNumber := os.Getenv("PORT")
	if len(portNumber) == 0 {
		// exit of the main programm
		log.Fatal("The PORT env is required")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(
		cors.Options{
			AllowedOrigins:   []string{"https://", "http://"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"*"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		},
	))

	v1Router := chi.NewRouter()
	// executes the http handler (handlerReadiness) when match with the endpoint pattern
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/error", handlerError)

	// split up into independent routers as V1Router
	router.Mount("/v1", v1Router)

	srv := http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%s", portNumber),
	}
	fmt.Println("Running on the port number: ", portNumber)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("something was wrong listening on the TCP network: %v", err)
	}
}
