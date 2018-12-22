#!/usr/bin/env python

import os
import RPi.GPIO as GPIO
import time
import signal
import requests

player_host = os.environ['PLAYER_HOST']

button_play = 16
button_volume_up = 18
button_volume_down = 13
button_track_next = 29
button_track_prev = 31

# cleanup the GPIOs when shutting down
class GracefulKiller:
  kill_now = False
  def __init__(self):
    signal.signal(signal.SIGINT, self.exit_gracefully)
    signal.signal(signal.SIGTERM, self.exit_gracefully)

  def exit_gracefully(self,signum, frame):
    self.kill_now = True

def setup():
    GPIO.setmode(GPIO.BOARD)
    GPIO.setup(button_play, GPIO.IN, pull_up_down=GPIO.PUD_UP)
    GPIO.setup(button_volume_up, GPIO.IN, pull_up_down=GPIO.PUD_UP)
    GPIO.setup(button_volume_down, GPIO.IN, pull_up_down=GPIO.PUD_UP)
    GPIO.setup(button_track_next, GPIO.IN, pull_up_down=GPIO.PUD_UP)
    GPIO.setup(button_track_prev, GPIO.IN, pull_up_down=GPIO.PUD_UP)
    print('Initialized GPIOs for buttons - ready to push')

def send_play_request():
    r = requests.get('http://' + player_host + '/rfid_player/play')
    print(r.status_code, r.reason)

def send_volume_request(change_event):
    r = requests.get('http://' + player_host + '/rfid_player/volume/' + change_event)
    print(r.status_code, r.reason)

def send_track_request(change_event):
    r = requests.get('http://' + player_host + '/rfid_player/track/' + change_event)
    print(r.status_code, r.reason)

def loop():
    killer = GracefulKiller()
    while True:
        button_play_state = GPIO.input(button_play)
        button_volume_up_state = GPIO.input(button_volume_up)
        button_volume_down_state = GPIO.input(button_volume_down)
        button_track_next_state = GPIO.input(button_track_next)
        button_track_prev_state = GPIO.input(button_track_prev)
        
        if button_play_state == False:
            print('Play/Pause button pressed...')
            send_play_request()
            while GPIO.input(button_play) == False:
                time.sleep(0.3)
        if button_volume_up_state == False:
            print('Volume up button pressed...')
            send_volume_request('up')
            while GPIO.input(button_volume_up) == False:
                time.sleep(0.3)
        if button_volume_down_state == False:
            print('Volume down button pressed...')
            send_volume_request('down')
            while GPIO.input(button_volume_down) == False:
                time.sleep(0.3)
        if button_track_next_state == False:
            print('Next track button pressed...')
            send_track_request('next')
            while GPIO.input(button_track_next) == False:
                time.sleep(0.3)
        if button_track_prev_state == False:
            print('Prev track button pressed...')
            send_track_request('prev')
            while GPIO.input(button_track_prev) == False:
                time.sleep(0.3)
    if killer.kill_now:
        print('SIGTERM detected')
        endprogram()

def endprogram():
    print('cleaning up the GPIOs')
    GPIO.cleanup()

if __name__ == '__main__':
    setup()
    loop()
