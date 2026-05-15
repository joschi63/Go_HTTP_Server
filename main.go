package main

import (
	"log"
	"net/http"

	"github.com/joschi64/Go_HTTP_Server/handler"
)

func main() {
	apiCfg := &handler.ApiConfig{}
	serveMux := http.NewServeMux()

	serveMux.Handle("/app/", http.StripPrefix("/app", apiCfg.MiddlewareMetricsInc(http.FileServer(http.Dir(".")))))
	serveMux.Handle("/app/assets/logo.png", http.FileServer(http.Dir("./assets/logo.png")))

	serveMux.HandleFunc("GET /api/healthz", handler.HandleHealthCheck)
	serveMux.HandleFunc("GET /admin/metrics", apiCfg.HandleServerHits)
	serveMux.HandleFunc("POST /admin/reset", apiCfg.HandleHitsReset)
	serveMux.HandleFunc("POST /api/validate_chirp", handler.HandleValidateChirp)

	server := &http.Server{
		Addr:    ":8080",
		Handler: serveMux,
	}

	log.Printf("Starting server on %s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed: %v", err)
	}
}
