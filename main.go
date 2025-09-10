package main

import (
	"log/slog"
	"net/http"
	"os"
	"proj4/internal/config"
	"proj4/internal/storage/sqlite"
	"proj4/internal/models"
	"proj4/internal/http_server/middleware/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	envLocal = "local"
	envDev = "dev"
	envProd = "prod"
)

func main() {



	// TODO config

	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)
	log.Info("Setup loger: ", slog.String("env", cfg.Env))
	log.Debug("debug message enabled")

	// TODO storage
	sqlite.InitDB(log, cfg.StoragePath)
	sqlite.Migrate(log, &models.User{})

	// sqlite.DB.Create(&models.User{Name: "Alex", Email: "Example"})
	
	// TODO route

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(logger.New(log))



	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})


	// TODO run
	http.ListenAndServe(":3000", router)

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}




