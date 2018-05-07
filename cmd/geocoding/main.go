package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/tetafro/geocoding/internal/places"
	"github.com/tetafro/geocoding/internal/router"
)

func main() {
	cfg := MustConfig()
	log := MustLogger(cfg.LogLevel, cfg.LogFormat)

	conn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s %s",
		cfg.PGHost, cfg.PGPort, cfg.PGDatabase,
		cfg.PGUsername, cfg.PGPassword, cfg.PGParams)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	placesRepo, err := places.NewPostgresRepo(db)
	if err != nil {
		log.Fatalf("Failed to init places repository: %v", err)
	}
	placesService := places.NewService(placesRepo)
	placesController := places.NewController(placesService, log)

	r := router.New(placesController)
	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Infof("Start listening at %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Errorf("Failed to start server: %v", err)
	}
}
