#!/bin/bash

# start mpv player in idle, set the json ipc socket path and fork the process
mpv --idle --no-video --input-ipc-server=/tmp/mpvsocket 2>/dev/null & 

# wait 2 seconds for mpv to start (we are running on a raspberry...)
sleep 2

/rpi-rfid-musicplayer_player