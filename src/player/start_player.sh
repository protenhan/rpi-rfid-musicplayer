#!/bin/bash

# start mpv player in idle, set the json ipc socket path and fork the process
mpv --idle --input-ipc-server=/tmp/mpvsocket & 

# wait 5 seconds for mpv to start (we are running on a raspberry...)
sleep 5

/rpi-rfid-musicplayer_player