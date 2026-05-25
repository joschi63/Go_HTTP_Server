package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/joschi64/Go_HTTP_Server/handler"
	"github.com/joschi64/Go_HTTP_Server/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	db, _ := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)

	apiCfg := &handler.ApiConfig{
		DB: dbQueries,
	}
	serveMux := http.NewServeMux()

	serveMux.Handle("/app/", http.StripPrefix("/app", apiCfg.MiddlewareMetricsInc(http.FileServer(http.Dir(".")))))
	serveMux.Handle("/app/assets/logo.png", http.FileServer(http.Dir("./assets/logo.png")))

	serveMux.HandleFunc("GET /api/healthz", handler.HandleHealthCheck)
	serveMux.HandleFunc("GET /admin/metrics", apiCfg.HandleServerHits)
	serveMux.HandleFunc("POST /admin/reset", apiCfg.HandleReset)
	serveMux.HandleFunc("POST /api/chirps", apiCfg.HandleChirp)
	serveMux.HandleFunc("GET /api/chirps", apiCfg.HandleChirps)
	serveMux.HandleFunc("GET /api/chirps/{chirpID}", apiCfg.GetChirp)
	serveMux.HandleFunc("POST /api/users", apiCfg.CreateUser)
	serveMux.HandleFunc("POST /api/login", apiCfg.HandleUserLogin)

	server := &http.Server{
		Addr:    ":8080",
		Handler: serveMux,
	}

	log.Printf("Starting server on %s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed: %v", err)
	}
}
