package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/nickchirgin/scpl/internal/playlist"
)
type Server struct {
	Router *mux.Router 
	Pl *playlist.Playlist
}

func NewServer(pl *playlist.Playlist) *Server {
	return &Server{Router: mux.NewRouter(), Pl: pl}
} 

func (s *Server) Run() {
	srv := &http.Server{
		Handler: s.Router,
		Addr: "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}