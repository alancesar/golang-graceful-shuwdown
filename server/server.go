package server

import (
	"context"
	"net/http"
)

type Server struct {
	*http.Server
}

func New(handler http.Handler, addr string) *Server {
	return &Server{
		Server: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
	}
}

func (h Server) Start() error {
	if err := h.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (h Server) Stop(ctx context.Context) error {
	return h.Shutdown(ctx)
}
