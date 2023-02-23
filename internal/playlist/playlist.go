package playlist

import (
	"container/list"
	"context"
	"fmt"
	"sync"
	"time"
)

type Playlist struct {
	 Tracks *list.List
	 current *list.Element
	 time int
	 mu sync.RWMutex
	 PlaySong chan Song 
	 Stop chan struct{}
	 nextprev chan bool  // true = next, false = prev
}

type Song struct {
	Title string
	Duration int
}

func NewPlaylist() Playlist {
	return Playlist{}
}

func (p *Playlist) Play(track Song) {
	p.mu.RLock()
	defer p.mu.RUnlock() 
	ctx := context.Background()
	_, cancel := context.WithTimeout(ctx, time.Duration(track.Duration - p.time))
	deadline, _ := ctx.Deadline()
	defer cancel() 
	select {
	case <-ctx.Done():
		fmt.Println("Song expired")
		p.Next()
	case data:=<-p.nextprev: 
		if data == true {
			p.Next()
		} else {
			p.Prev()
		}
	case <-p.Stop:
		fmt.Println(time.Until(deadline))
	}	
}

func (p *Playlist) Pause() {
	p.Stop <- struct{}{}
}

func (p *Playlist) AddSong(track Song) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.Tracks.PushBack(track)
	p.PlaySong <- track
}

func (p *Playlist) Next() {
	p.mu.RLock()
	defer p.mu.RUnlock()
	p.time = 0
	if p.current.Next().Value != nil {
		p.current.Value = p.current.Next().Value
		p.Play(p.current.Value.(Song))
	}
}

func (p *Playlist) Prev() {
	p.mu.RLock()
	defer p.mu.RUnlock()
	p.time = 0
	if p.current.Next().Value != nil {
		p.current.Value = p.current.Prev().Value
		p.Play(p.current.Value.(Song))
	}
}