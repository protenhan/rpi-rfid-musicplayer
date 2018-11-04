#!/usr/bin/env python

import threading

from rfid_reader import RFIDReader
from button_control import ButtonController


if __name__ == '__main__':
    rfid_thread = threading.Thread(RFIDReader.activate_reader())
    button_thread = threading.Thread(ButtonController.activate_buttons())