package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/IskanderSh/vk-task/internal/app"
	"github.com/IskanderSh/vk-task/internal/config"
)

const (
	envLocal = "local"
	envProd  = "prod"

	debugLvl = "DEBUG"
	infoLvl  = "INFO"
	warnLvl  = "WARN"
	errorLvl = "ERROR"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg)
	log.Info("logger initialized successfully")

	application := app.NewServer(log, cfg)
	go func() {
		if err := application.HTTPServer.ListenAndServe(); err != nil {
			log.Error(err.Error())
		}
	}()
	log.Info("application started successfully")

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop
	log.Info("stopping application", slog.String("signal", sign.String()))
	application.HTTPServer.Shutdown(context.Background())

	log.Info("application stopped")
}

func setupLogger(cfg *config.Config) *slog.Logger {
	var log *slog.Logger

	switch cfg.Env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: getLogLevel(cfg.LogLevel)}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: getLogLevel(cfg.LogLevel)}))
	}

	return log
}

func getLogLevel(lvl string) slog.Level {
	var res slog.Level

	switch strings.ToUpper(lvl) {
	case debugLvl:
		res = slog.LevelDebug
	case infoLvl:
		res = slog.LevelInfo
	case warnLvl:
		res = slog.LevelWarn
	case errorLvl:
		res = slog.LevelError
	}

	return res
}
