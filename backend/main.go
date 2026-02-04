package main

import (
	"backend/internal/handlers"
	"backend/internal/r2"
	"backend/internal/utils"
	"log"
	"net/http"
)

func main() {
	envCfg := utils.SetupEnvCfg()

	// Setup r2 integration
	r2, err := r2.Setup()
	if err != nil {
		log.Fatal(err)
	}

	var recentBuilds []handlers.RecentBuild

	cfg := &handlers.APIConfig{
		Env:          envCfg,
		R2:           r2,
		RecentBuilds: recentBuilds,
	}

	mux := http.NewServeMux()

	// System
	mux.HandleFunc("/api/health", cfg.Health)

	// Builds
	mux.HandleFunc("GET /api/getBuild/{id}", cfg.GetBuild)
	mux.HandleFunc("POST /api/newBuild", cfg.NewBuild)

	mux.HandleFunc("GET /api/recent", cfg.GetRecentBuilds)

	srv := &http.Server{
		Addr:    ":" + cfg.Env.Port,
		Handler: mux,
	}

	log.Printf("Serving on: http://localhost:%s/\n", cfg.Env.Port)
	log.Fatal(srv.ListenAndServe())
}
