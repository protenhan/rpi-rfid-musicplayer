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
		event, _ := device.Read()
		for i, v := range event {
			fmt.Println("A event was triggered: " + "index: " + i + " and value: " + v)
		}
	}
}

func init() {
	devicePath = os.Getenv("RFID_DEVICE_PATH")
}