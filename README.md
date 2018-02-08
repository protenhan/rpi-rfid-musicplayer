[![Build Status](https://travis-ci.org/protenhan/rpi-rfid-musicplayer.svg?branch=master)](https://travis-ci.org/protenhan/rpi-rfid-musicplayer)
[![](https://images.microbadger.com/badges/image/protenhan/rpi-rfid-musicplayer.svg)](https://microbadger.com/images/protenhan/rpi-rfid-musicplayer "Get your own image badge on microbadger.com")
[![](https://images.microbadger.com/badges/version/protenhan/rpi-rfid-musicplayer.svg)](https://microbadger.com/images/protenhan/rpi-rfid-musicplayer "Get your own version badge on microbadger.com")

# rpi-rfid-musicplayer
Musicplayer for the rapsberry pi that plays music based on RFID cards that are presented to a RFID reader.


## Requirements

* Raspberry Pi 3 (a Raspberry Pi 2 will propably also work)
* USB RFID card reader (I'm using [this cheap china one from amazon](https://www.amazon.de/gp/product/B00HSDOTTU/ref=oh_aui_detailpage_o02_s00?ie=UTF8&psc=1))
                       ![Card reader image](https://images-na.ssl-images-amazon.com/images/I/51GC53JXfPL._SX425_.jpg)

## Running Musicplayer

Set the RFID_DEVICE_PATH environment variable to specify the devicePath of your RFID reader

1. Get the id of your RFID reader
```bash
user@raspi-soundbox:~ $ ls /dev/input/by-id/
usb-Sycreader_USB_Reader_08FF20150112-event-kbd
```
2. Set the Device Path to your environment
```bash
export RFID_DEVICE_PATH=/dev/input/by-id/usb-Sycreader_USB_Reader_08FF20150112-event-kbd
```
3. Run the docker container
```bash
docker run --rm -e RFID_DEVICE_PATH=$RFID_DEVICE_PATH --device=$RFID_DEVICE_PATH protenhan/rpi-rfid-musicplayer
```
