package web

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	log.Printf("Server runs on http://localhost%s\n", s.httpServer.Addr)
	err := s.httpServer.ListenAndServe()
	return fmt.Errorf("Server.Run: %w", err)
}
