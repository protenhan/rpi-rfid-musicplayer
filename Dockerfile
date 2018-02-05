FROM hypriot/rpi-alpine
MAINTAINER @protenhan

RUN apk update &&\
    apk -U add \
        python3

COPY src/python /rfid-musicplayer/

RUN pip install -r /rfid-musicplayer/requirements.txt

CMD /rfid-musicplayer/rfid_input_reader.py
