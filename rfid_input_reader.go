package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var (
	device string
)

func main() {
	fmt.Println("device: ", device)

	for {
		cardID, err := ioutil.ReadFile(device)
		if err != nil {
			panic(err)
		}

		fmt.Println(cardID)
	}

}

func init() {
	device = os.Getenv("RFID_HID_DEVICE")
}
