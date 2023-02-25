package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/nickchirgin/scpl/internal/playlist"
	"github.com/nickchirgin/scpl/internal/playlist/handlers"
)
type Server struct {
	Router *mux.Router 
	Handler *handlers.Handler 
}

func NewServer(pl *playlist.Playlist) *Server {
	return &Server{Router: mux.NewRouter(), Handler: &handlers.Handler{Pl: pl}}
} 

func (s *Server) Run() {
	srv := &http.Server{
		Handler: s.Router,
		Addr: ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
	fmt.Println("Server is running")
}