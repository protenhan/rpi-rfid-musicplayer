package main

import (
	"encoding/json"
	"log"
	"net/http"

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

// our main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/rfid_player/playlist/{id}", startPlaybackOfPlaylistWithID).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func startPlaybackOfPlaylistWithID(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	var Playlist playlist
	_ = json.NewDecoder(request.Body).Decode(&Playlist)
	Playlist.ID = params["id"]
	currentPlaylist = Playlist
	json.NewEncoder(writer).Encode(currentPlaylist)
}
