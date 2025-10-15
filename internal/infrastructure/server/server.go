package server

import (
	"fmt"
	"log/slog"
	"net/http"
)

type Server struct {
	addr   string
	mux    *http.ServeMux
	logger *slog.Logger
}

func NewServer(port int, handler Handler, logger *slog.Logger) Server {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello-world", handler.getHello)

	return Server{
		addr:   fmt.Sprintf(":%d", port),
		mux:    mux,
		logger: logger,
	}
}

func (s Server) Start() error {
	s.logger.Info("Starting server", slog.String("addr", s.addr))
	return http.ListenAndServe(s.addr, s.mux)
}
