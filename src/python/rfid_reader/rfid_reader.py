#!/usr/bin/env python

import os
import subprocess

from evdev import InputDevice
from select import select


class RFIDReader:
    def play_folder(self, folder_name):
        folder_path = "/audio/" + folder_name + ".mp3"
        subprocess.check_output(['play', folder_path])

    def activate_reader(self):
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