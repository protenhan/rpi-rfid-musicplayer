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
	for true {
		event := device.Read()
		fmt.Println("A event was triggered: " + event)
	}
}

func init() {
	devicePath = os.Getenv("RFID_DEVICE_PATH")
}