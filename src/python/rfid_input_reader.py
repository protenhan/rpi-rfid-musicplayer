import string
import os

from evdev import InputDevice
from select import select

devicePath = os.environ['RFID_DEVICE_PATH']

keys = "X^1234567890XXXXqwertzuiopXXXXasdfghjklXXXXXyxcvbnmXXXXXXXXXXXXXXXXXXXXXXX"
dev = InputDevice(devicePath)

code = ""

while True:
    r,w,x = select([dev], [], [])
    for event in dev.read():
        if event.type==1 and event.value==1:
            character = keys[ event.code ]
            if character is "X":
                print(code)
                code = ""
            else:
                code += character

