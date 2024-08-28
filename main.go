package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dagregi/blog-aggregator/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT enviromental variable not set")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL enviromental variable not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	cfg := apiConfig{
		DB: dbQueries,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/users", cfg.createUserHandler)
	mux.HandleFunc("GET /v1/users", cfg.middlewareAuth(cfg.getUserHandler))

	mux.HandleFunc("POST /v1/feeds", cfg.middlewareAuth(cfg.createFeedHandler))
	mux.HandleFunc("GET /v1/feeds", cfg.getFeedsHandler)

	mux.HandleFunc("POST /v1/feed_follows", cfg.middlewareAuth(cfg.createFeedFollowHandler))
	mux.HandleFunc("GET /v1/feed_follows", cfg.middlewareAuth(cfg.getFeedFollowsHandler))
	mux.HandleFunc("DELETE /v1/feed_follows/{feedFollowID}", cfg.middlewareAuth(cfg.deleteFeedFollowHandler))

	mux.HandleFunc("GET /v1/healthz", healthzHandler)
	mux.HandleFunc("GET /v1/err", errHandler)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	const collectionConcurrency = 10
	const collectionInterval = 5 * time.Minute
	go startScraping(dbQueries, collectionConcurrency, collectionInterval)

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}
