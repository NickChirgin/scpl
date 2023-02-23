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
	 Current *list.Element
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
	return Playlist{Tracks: list.New(), Current: &list.Element{}, PlaySong: make(chan Song), Stop: make(chan struct{}), nextprev: make(chan bool), time: 0}
}

func (p *Playlist) Play(track Song) {
	p.mu.RLock()
	defer p.mu.RUnlock() 
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration((track.Duration - p.time)) * time.Second)
	defer cancel() 
	deadline, _ := ctx.Deadline()
	p.Current.Value = track
	fmt.Println(p.Current.Value)
	select {
	case <-ctx.Done():
		fmt.Println("Song expired", p.Tracks.Len())
		p.PlaySong <- p.Current.Value.(Song)
		if p.Current.Next() != nil {
			p.Next()
		}
	case data:=<-p.nextprev: 
		if data == true {
			p.Next()
		} else {
			p.Prev()
		}
	case <-p.Stop:
		fmt.Println(time.Until(deadline))
	default:
		fmt.Println("TUTU")
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
	fmt.Println("ESt kto")
	p.time = 0
	if p.Current.Next() != nil || p.Current.Next().Value != (Song{}) {
		p.Current.Value = p.Current.Next().Value
		p.Play(p.Current.Value.(Song))
	}
}

func (p *Playlist) Prev() {
	p.mu.RLock()
	defer p.mu.RUnlock()
	p.time = 0
	if p.Current.Next().Value != nil {
		p.Current.Value = p.Current.Prev().Value
		p.Play(p.Current.Value.(Song))
	}
}