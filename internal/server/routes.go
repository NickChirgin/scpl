package server

import (
	"net/http"
)

func (s *Server) RegisterRoutes() {
	s.Router.HandleFunc("/play", s.Handler.Play).Methods(http.MethodGet)
	s.Router.HandleFunc("/", s.Handler.Base).Methods(http.MethodGet)
	s.Router.HandleFunc("/pause", s.Handler.Pause).Methods(http.MethodGet)
	s.Router.HandleFunc("/next", s.Handler.Next).Methods(http.MethodGet)
	s.Router.HandleFunc("/prev", s.Handler.Prev).Methods(http.MethodGet)
	s.Router.HandleFunc("/addsong", s.Handler.AddSong).Methods(http.MethodPost)
}