#!/usr/bin/env python

import RPi.GPIO as GPIO
import time
import signal

button_play = 16

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
    print('Initialized GPIOs for buttons - ready to push')

def loop():
    killer = GracefulKiller()
    while True:
        button_state = GPIO.input(button_play)
        if button_state == False:
            # TODO: send web request to player
            print('Play/Pause button pressed...')
            while GPIO.input(button_play) == False:
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
