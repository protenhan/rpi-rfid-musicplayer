FROM hypriot/rpi-alpine
MAINTAINER @protenhan

RUN apk update &&\
    apk -U add \
        python3 \
        linux-headers \
        python3-dev \
        gcc \
        musl-dev

COPY src/python /rfid-musicplayer/

RUN pip3 install -r /rfid-musicplayer/requirements.txt

CMD python3 /rfid-musicplayer/rfid_input_reader.py
