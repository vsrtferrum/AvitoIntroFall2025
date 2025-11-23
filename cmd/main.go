package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/vsrtferrum/AvitoIntroFall2025/internal/aplication"
	"github.com/vsrtferrum/AvitoIntroFall2025/internal/infrastructure/api"
	"github.com/vsrtferrum/AvitoIntroFall2025/internal/infrastructure/config"
	httpGen "github.com/vsrtferrum/AvitoIntroFall2025/internal/infrastructure/http"
	"github.com/vsrtferrum/AvitoIntroFall2025/internal/infrastructure/storage"
	"github.com/vsrtferrum/AvitoIntroFall2025/pkg/logger"
)

func main() {

	logger := logger.NewLogger()
	err := logger.Raise()
	if err != nil {
		panic(err)
	}

	cfg, err := config.ReadConfig(logger)
	if err != nil {
		panic(err)
	}

	db, err := storage.NewStorage(context.Background(), cfg, logger)
	if err != nil {
		panic(err)
	}

	users, teams, prs, err := db.GetAllData()
	if err != nil {
		panic(err)
	}

	app := aplication.NewGitModel(users, teams, prs, db, logger)

	apiHandler := api.NewAPIHandler(app, logger)

	router := chi.NewRouter()

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	})

	httpHandler := httpGen.HandlerFromMux(apiHandler, router)

	server := &http.Server{
		Addr:    ":" + "8080",
		Handler: httpHandler,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {

		panic(err)
	}
}
