package main

import (
	"encoding/json"
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
var mpcClient = mpv.NewClient(mpv.NewIPCClient(jsonIpcSocketPath))

// our main function
func main() {
	println("connecting to mpv")

	playerState = STOPPED
	println("player started and listening")
	router := mux.NewRouter()
	router.HandleFunc("/rfid_player/playlist/{id}", startPlaybackOfPlaylistWithID).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func startPlaybackOfPlaylistWithID(writer http.ResponseWriter, request *http.Request) {
	println("recieved POST request")
	params := mux.Vars(request)
	var Playlist playlist
	_ = json.NewDecoder(request.Body).Decode(&Playlist)
	Playlist.ID = params["id"]
	println("New Playlist ID is " + Playlist.ID)
	currentPlaylist = Playlist
	json.NewEncoder(writer).Encode(currentPlaylist)
	startPlayback()
}

func startPlayback() {
	playlistPath := "/audio/" + currentPlaylist.ID + "/" + currentPlaylist.ID + ".m3u"
	// this switch statement is here only as placeholder, if I ever find a reason to implement a diefferent behaviour.
	mpcClient.LoadList(playlistPath, "LoadListModeReplace")
	println("loaded playlist " + playlistPath + " in MPV")
	resumePlayback()
}

func resumePlayback() {
	mpcClient.SetPause(false)
	playerState = PLAYING
	println("resumed playback")
}
