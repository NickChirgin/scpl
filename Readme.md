### Module for working with a playlist

#### The design of API was described in the test task.

##### Available methods

    Play - starts playing
    Pause - pauses playback
    AddSong - adds a song to the end of the playlist
    Next - plays the next song
    Prev - plays the previous song

The methods work in separate goroutines.

### API for interacting with the playlist

* GET /play - starts playback from the beginning/continues from the last place.
* GET /pause - stops playback at the current timer.
* GET /prev - starts playing the previous track.
* GET /next - starts playing the next track.
* POST /addsong - adds a track to the end of the playlist. Fields required for successful addition: Title - string, Duration - int

### An example of how it works is shown in this gif:

![How it works](https://raw.githubusercontent.com/NickChirgin/scpl/master/gi.gif)

### How to run?

The server is launched on port 8080.
```
git clone https://github.com/NickChirgin/scpl.git
```
```
make run
```
