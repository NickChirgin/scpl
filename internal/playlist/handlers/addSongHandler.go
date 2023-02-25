package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nickchirgin/scpl/internal/playlist"
)

func(h *Handler) AddSong(w http.ResponseWriter, r *http.Request){
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()
	song := playlist.Song{}  
	err := dec.Decode(&song)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	fmt.Println(song)
	h.Pl.Add <- song
}