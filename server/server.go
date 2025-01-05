package server

import (
	"fmt"
	"net/http"

	"github.com/climinal/server/config"
)

type Server struct {
	config *config.Config
}

func New(cfg *config.Config) *Server {
	return &Server{
		config: cfg,
	}
}

func (s *Server) Start() error {
	// TODO: Implement WebSocket handler and message routing
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Climinal Server Running")
	})

	return http.ListenAndServe(fmt.Sprintf(":%d", s.config.Port), nil)
}
