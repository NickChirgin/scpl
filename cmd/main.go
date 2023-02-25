package main

import (
	"sync"
	"time"

	"github.com/nickchirgin/scpl/internal/playlist"
	"github.com/nickchirgin/scpl/internal/server"
)

func main() {
	pl := playlist.NewPlaylist()
	var wg sync.WaitGroup
	pl.Tracks.PushBack(playlist.Song{Title: "Big Stan", Duration: 7})
	pl.Tracks.PushBack(playlist.Song{Title: "Mockingbird", Duration: 10})
	pl.Tracks.PushBack(playlist.Song{Title: "Venom", Duration: 15})
	pl.Current = pl.Tracks.Front()
	go pl.Reciever()
	go pl.Play()
	wg.Add(2)
	go func(){
		time.Sleep(4 * time.Second)
		pl.Stop <- struct{}{}
		pl.Add <- playlist.Song{Title: "Rap God", Duration: 9} 
		time.Sleep(2 * time.Second)
		pl.PlaySong <- struct{}{}
		time.Sleep(7 * time.Second)
		pl.Nextprev <- true
		wg.Done()
	}()
	wg.Wait()
	s := server.NewServer(&pl)
	s.Run()
}