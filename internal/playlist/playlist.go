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
	 Add chan Song 
	 Nextprev chan bool  // true = next, false = prev
	 PlaySong chan struct{} 
	 Stop chan struct{}
}

type Song struct {
	Title string
	Duration int
}

func NewPlaylist() Playlist {
	return Playlist{Tracks: list.New(), Current: &list.Element{}, PlaySong: make(chan struct{}), Stop: make(chan struct{}), Nextprev: make(chan bool), time: 0, Add: make(chan Song)}
}

func (p *Playlist) Play() {
	p.mu.RLock()
	defer p.mu.RUnlock() 
	if p.Tracks.Len() == 0 {
		fmt.Println("Playlist is empty")
		return
	}
	if p.Current == nil {
		p.Current = p.Tracks.Front()
	}
	if p.time ==  0{
		p.time = p.Current.Value.(Song).Duration
	}
	fmt.Println(p.time)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(p.time) * time.Second)
	defer cancel() 
	deadline, _ := ctx.Deadline()
	fmt.Println(p.Current.Next(), ctx)
	select {
	case <-ctx.Done():
		fmt.Printf("%s song has ended", p.Current.Value.(Song).Title)
		p.Next()
	case <-p.Stop: 
		p.time = int(time.Until(deadline).Seconds())
		fmt.Println(p.time)
		break
	}	
}

func (p *Playlist) Pause() {
	p.Stop <- struct{}{}
}

func (p *Playlist) AddSong(track Song) {
	p.mu.Lock()
	defer p.mu.Unlock()
	fmt.Printf("%s song has been added", track.Title)
	p.Tracks.PushBack(track)
}

func (p *Playlist) Next() {
	p.mu.RLock()
	defer p.mu.RUnlock()
	p.time = 0
	if p.Current.Next() != nil || p.Current.Next().Value != (Song{}) {
		p.Current = p.Current.Next()
		p.Play()
	}
}

func (p *Playlist) Prev() {
	p.mu.RLock()
	defer p.mu.RUnlock()
	p.time = 0
	if p.Current.Prev() != nil {
		p.Current = p.Current.Prev()
		p.Play()
	}
}