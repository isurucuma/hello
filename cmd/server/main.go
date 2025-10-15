package main

import (
	"github.com/isurucuma/hello/internal/infrastructure/server"
	"github.com/isurucuma/hello/internal/usecase"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	helloUsecase := usecase.NewHelloImpl(logger)
	helloHandler := server.NewHandler(helloUsecase, logger)
	srv := server.NewServer(8080, helloHandler, logger)
	if err := srv.Start(); err != nil {
		logger.Error("Failed to start srv", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
