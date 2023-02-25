package server

import (
	"net/http"

	"github.com/nickchirgin/scpl/internal/playlist/handlers"
)

func (s *Server) RegisterRoutes() {
	s.Router.HandleFunc("/play", handlers.Play).Methods(http.MethodGet)
	s.Router.HandleFunc("/pause", handlers.Pause).Methods(http.MethodGet)
	s.Router.HandleFunc("/next", handlers.Next).Methods(http.MethodGet)
	s.Router.HandleFunc("/prev", handlers.Prev).Methods(http.MethodGet)
	s.Router.HandleFunc("/addsong", handlers.AddSong).Methods(http.MethodPost)
}