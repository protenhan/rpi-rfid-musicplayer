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
	fmt.Println("device path: ", devicePath)
	fmt.Println(evdev.IsInputDevice(devicePath))
	device, _ := evdev.Open(devicePath)
	fmt.Println(device)
}

func init() {
	devicePath = os.Getenv("RFID_DEVICE_PATH")
}