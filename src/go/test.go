package main

import (
	"fmt"
	"os"
)

var (
	devicePath string
)

func main() {
	devicePath = os.Getenv("RFID_DEVICE_PATH")
	fmt.Print("device path: ", devicePath)
}
