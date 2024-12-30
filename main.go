package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/daniela2001-png/rss_aggregator_project/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // pq is a pure Go Postgres driver for the database/sql package.
)

type apiConf struct {
	DB *database.Queries
}

func main() {

	// load envs from .env file
	godotenv.Load(".env")
	portNumber := os.Getenv("PORT")
	if len(portNumber) == 0 {
		// exit of the main programm
		log.Fatal("The PORT env is required")
	}

	dbURL := os.Getenv("DB_URL")
	if len(dbURL) == 0 {
		// exit of the main programm
		log.Fatal("The DB_URL env is required")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can not connect to database")
	}
	dbQuery := database.New(conn)
	apiCnf := apiConf{
		DB: dbQuery,
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

	// handlerCreateUser works as a method of apiConf struct
	v1Router.Post("/users", apiCnf.handlerCreateUser)

	// handlerGetUser works as a method of apiConf struct
	// Calls middlewareAuth before we get user information
	v1Router.Get("/user", apiCnf.middlewareAuth(apiCnf.handlerGetUser))

	// handlerCreateFeed works as a method of apiConf struct
	v1Router.Post("/feeds", apiCnf.middlewareAuth(apiCnf.handlerCreateFeed))
	v1Router.Get("/feeds", apiCnf.handlerGetFeeds)

	v1Router.Post("/feed_follows", apiCnf.middlewareAuth(apiCnf.handlerCreateFeedFollow))
	v1Router.Get("/feed_follows", apiCnf.middlewareAuth(apiCnf.handlerGetListOfFeedsOfAnUser))
	v1Router.Delete("/feed_follows/{feedFollowID}", apiCnf.middlewareAuth(apiCnf.handlerUnFollowFeedID))

	// split up into independent routers as V1Router
	router.Mount("/v1", v1Router)

	srv := http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%s", portNumber),
	}
	fmt.Println("Running on the port number: ", portNumber)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("something was wrong listening on the TCP network: %v", err)
	}
}
