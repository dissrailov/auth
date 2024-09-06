package main

import (
	"auth/internal/config"
	"auth/internal/handlers"
	"auth/internal/lib/logger/handlers/slogpretty"
	"auth/internal/lib/logger/sl"
	"auth/internal/repo"
	"auth/internal/repo/postgres"
	"auth/internal/server"
	"auth/internal/service"
	"fmt"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envProd  = "prod"
	envDev   = "dev"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting auth service")
	fmt.Println(cfg.Database)
	db, err := postgres.NewPostgresDB(cfg.Database)
	if err != nil {
		log.Error("failed to connect to database", sl.Err(err))
		os.Exit(1)
	}
	fmt.Println(cfg.HTTPServer)
	repo := repo.NewRepository(db)
	service := service.NewService(repo)
	handlers := handlers.NewHandler(service, cfg)
	srv := server.New(&cfg.HTTPServer, handlers.InitRoutes())
	if err := srv.Run(); err != nil {
		log.Error("failed to start server", sl.Err(err))
		os.Exit(1)
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = setupPrettySlog()

	case envProd:
		log = setupPrettySlog()
	}
	return log
}
func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
