package main

import (
	"fmt"

	"github.com/nickchirgin/scpl/internal/playlist"
)

func main() {
	pl := playlist.NewPlaylist()
	pl.Tracks.PushBack(playlist.Song{Title: "Yo", Duration: 10})
	pl.Tracks.PushBack(playlist.Song{Title: "Demolisher", Duration: 11})
	go pl.AddSong(playlist.Song{Title: "Hey", Duration: 15})
	go pl.Play(pl.Tracks.Back().Value.(playlist.Song))
	select {
	case data:=<-pl.PlaySong:
		fmt.Println(data)
	}
}