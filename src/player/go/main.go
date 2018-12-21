package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/blang/mpv"
	"github.com/gorilla/mux"
)

type state int

// Constant defining the current player state
const (
	STOPPED state = iota
	PLAYING
	PAUSED
)

type playerStatus struct {
	PlayerState     state `json:"player_state,omitempty"`
	CurrentPlaylist playlist
	Volume          int16 `json:"current_volume,omitempty"`
}

type playlist struct {
	ID string `json:"id,omitempty"`
}

var currentPlaylist playlist
var currentVolume int16
var playerState state

var jsonIpcSocketPath = "/tmp/mpvsocket"
var mpvClient = mpv.NewClient(mpv.NewIPCClient(jsonIpcSocketPath))

// our main function
func main() {
	playerState = STOPPED
	fmt.Println("player started and listening")
	router := mux.NewRouter()
	router.HandleFunc("/rfid_player/playlist/{id}", startPlaybackOfPlaylistWithID).Methods("POST")
	router.HandleFunc("/rfid_player/play", handlePlayRequest).Methods("GET")
	router.HandleFunc("/rfid_player/volume/up", handelVolumeUpRequest).Methods("GET")
	router.HandleFunc("/rfid_player/volume/down", handelVolumeDownRequest).Methods("GET")
	router.HandleFunc("/rfid_player/track/next", handelNextTrackRequest).Methods("GET")
	router.HandleFunc("/rfid_player/track/prev", handelPrevTrackRequest).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func startPlaybackOfPlaylistWithID(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("recieved POST request")
	params := mux.Vars(request)
	var Playlist playlist
	_ = json.NewDecoder(request.Body).Decode(&Playlist)
	Playlist.ID = params["id"]
	fmt.Println("New Playlist ID is " + Playlist.ID)
	currentPlaylist = Playlist
	json.NewEncoder(writer).Encode(currentPlaylist)
	startPlayback()
}

func handlePlayRequest(writer http.ResponseWriter, request *http.Request) {
	if playerState == PLAYING {
		pausePlayback()
	} else if playerState == PAUSED {
		resumePlayback()
	}
}

func handelVolumeDownRequest(writer http.ResponseWriter, request *http.Request) {
	currentVolume, _ := mpvClient.Volume()
	if currentVolume > 0 {
		fmt.Printf("Current Volume: %.2f\n", currentVolume)
		newVolume := currentVolume - float64(10.0)
		mpvClient.SetProperty("volume", newVolume)
		fmt.Printf("New Volume: %.2f\n", newVolume)
	}
}

func handelVolumeUpRequest(writer http.ResponseWriter, request *http.Request) {
	currentVolume, _ := mpvClient.Volume()
	if currentVolume < 100 {
		fmt.Printf("Current Volume: %.2f\n", currentVolume)
		newVolume := currentVolume + float64(10.0)
		mpvClient.SetProperty("volume", newVolume)
		fmt.Printf("New Volume: %.2f\n", newVolume)
	}
}

func handelNextTrackRequest(writer http.ResponseWriter, request *http.Request) {

}

func handelPrevTrackRequest(writer http.ResponseWriter, request *http.Request) {

}

func startPlayback() {
	playlistPath := "/audio/" + currentPlaylist.ID + "/" + currentPlaylist.ID + ".m3u"
	// this switch statement is here only as placeholder, if I ever find a reason to implement a diefferent behaviour.
	mpvClient.LoadList(playlistPath, "replace")
	fmt.Println("loaded playlist " + playlistPath + " in MPV")
	resumePlayback()
}

func resumePlayback() {
	mpvClient.SetPause(false)
	playerState = PLAYING
	fmt.Println("resumed playback")
}

func pausePlayback() {
	mpvClient.SetPause(true)
	playerState = PAUSED
	fmt.Println("paused playback")
}
