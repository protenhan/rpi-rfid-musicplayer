[![Build Status](https://travis-ci.org/protenhan/rpi-rfid-musicplayer.svg?branch=master)](https://travis-ci.org/protenhan/rpi-rfid-musicplayer)
[![](https://images.microbadger.com/badges/image/protenhan/rpi-rfid-musicplayer.svg)](https://microbadger.com/images/protenhan/rpi-rfid-musicplayer "Get your own image badge on microbadger.com")
[![](https://images.microbadger.com/badges/version/protenhan/rpi-rfid-musicplayer.svg)](https://microbadger.com/images/protenhan/rpi-rfid-musicplayer "Get your own version badge on microbadger.com")

# rpi-rfid-musicplayer
Musicplayer for the rapsberry pi that plays music based on RFID cards that are presented to a RFID reader.

## Running Musicplayer

Set the RFID_DEVICE_PATH environment variable to specify the devicePath of your RFID reader

```
docker run --rm -e RFID_DEVICE_PATH=/dev/input/by-id/usb-HID_OMNIKEY_5127_CK_01010053423438303000835748112531-event-kbd --device=/dev/input/by-id/usb-HID_OMNIKEY_5127_CK_01010053423438303000835748112531-event-kbd protenhan/rpi-rfid-musicplayer
```
