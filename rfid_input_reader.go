package main

import (
"fmt"
"os"
"github.com/gvalkov/golang-evdev"
)

var (
	devicePath string
)

func main() {
	fmt.Print("device path: ", devicePath)
	fmt.Print(evdev.IsInputDevice(devicePath))
}

func init() {
	devicePath = os.Getenv("RFID_DEVICE_PATH")
}