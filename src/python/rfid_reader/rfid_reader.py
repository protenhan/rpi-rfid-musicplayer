#!/usr/bin/env python
import os

from evdev import InputDevice
from select import select
import requests


player_host = os.environ['PLAYER_HOST']
def send_playlist_cmd(self, code):
    r = requests.post('http://' + player_host + '/rfid_player/playlist/' + code)
    print(r.status_code, r.reason)

if __name__ == "__main__":
    devicePath = os.environ['RFID_DEVICE_PATH']

    keys = "X^1234567890XXXXqwertzuiopXXXXasdfghjklXXXXXyxcvbnmXXXXXXXXXXXXXXXXXXXXXXX"
    dev = InputDevice(devicePath)

    code = ""

    while True:
        r, w, x = select([dev], [], [])
        for event in dev.read():
            if event.type == 1 and event.value == 1:
                character = keys[event.code]
                if character is "X":
                    print(code + " was read")
                    self.play_folder(code)
                    code = ""
                else:
                    code += character
