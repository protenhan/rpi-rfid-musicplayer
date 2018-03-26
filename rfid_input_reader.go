package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var (
	devicePath string
)

func main() {
	fmt.Println("device path: ", devicePath)

	for {
		cardID, err := ioutil.ReadFile(devicePath)
		if err != nil {
			panic(err)
		}

		fmt.Println(cardID)
	}

}

func init() {
	devicePath = os.Getenv("RFID_HID_DEVICE_PATH")
}
