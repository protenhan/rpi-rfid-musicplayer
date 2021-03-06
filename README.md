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

1. Flash a [current Raspbian image](https://www.raspberrypi.org/downloads/raspbian/) onto your SD card
2. Enable the uart GPIO pin so we can add a power LED without coding later on and set the `gpu_mem` to 32Mb or less in */boot/config.txt* . Who needs UI anyway? 😜 
    ```
    gpu_mem=32
    enable_uart=1
    ```
3. **(optional)** Configure your Wifi in */boot/wpa_supplicant.conf* 
```
ctrl_interface=DIR=/var/run/wpa_supplicant GROUP=netdev
update_config=1
country=«your_ISO-3166-1_two-letter_country_code» [US|DE|...]

network={
    ssid="«your_SSID»"
    psk="«your_PSK»"
    key_mgmt=WPA-PSK
}
```
4. Enable SSH by putting an empty file named `ssh` onto */boot*
5. Put SD Card into Raspi and boot it. 
6. Finish the configuration of your Raspi with `raspi-config` and change the default password with `passwd`
7. Install Docker-CE [official documentation](https://docs.docker.com/install/linux/docker-ce/debian/)
8. Copy audio files onto the raspberry (e.g. with scp)
    1. the player will look for folders that have the same name as the RFID Card Tag (e.g. 0015439814). Therefore every album is in one folder and gets named accordingly
    2. Run the very helpful [playlist creation script](https://gist.github.com/scarlson/944860) inside your audio root folder. It will create .m3u playlists for each folder that are named as the contianing folder. These playlist files are then passed to MPV to play the content of the folders.

My RFID Reader presents itself as an HumanInterfaceDevice (a.k.a. keyboard) to the system. Set the RFID_DEVICE_PATH environment variable to specify the devicePath of your RFID reader  

```
docker run --rm -e RFID_DEVICE_PATH=/dev/input/by-id/usb-Sycreader_USB_Reader_08FF20150112-event-kbd --device /dev/input/by-id/usb-Sycreader_USB_Reader_08FF20150112-event-kbd --device /dev/snd -v /etc/asound.conf:/etc/asound.conf protenhan/rpi-rfid-musicplayer
```


## Architecture overview
(And before you ask, of course this is overengineered 👨🏻‍💻)

The whole player experience consists of 3 containers. 

### The RFID_READER container 
Attaches to the USB RFID reader, captures the event, when a RFID card is read and then sends a POST request to the player container with the id of the RFID card.

### The PLAYER container
This container features the actual audio playback logic. Playing the audio files is handled by [MPV Player](https://mpv.io) which is running headless. MPV player is controlled by a small go application that runs REST endpoints for POSTing RFID card IDs and control events (pause, volume up, etc) and then forwards these events to the MPV player instance. 

## Wiring Layout

The player features buttons for 
* Play/pause
* Volume up
* Volume down
* next track ⏭
* previous track ⏮

The wiring layout for the Raspberry Pi should look like this ![Breadbord layout](docs/img/rfid_player_wiring_bb.png)