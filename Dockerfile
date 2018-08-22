FROM hypriot/rpi-alpine
MAINTAINER @protenhan

RUN apk update &&\
    apk -U add \
        python3 \
        linux-headers \
        python3-dev \
        gcc \
        musl-dev \
        alsa-utils \
        alsa-utils-doc \
        alsa-lib \
        alsaconf \
        sox

# Configure alsa audio
RUN sed '/audio:x:18/s/$/root/' /etc/group


COPY src/python /rfid-musicplayer/

RUN pip3 install -r /rfid-musicplayer/requirements.txt

CMD python3 /rfid-musicplayer/rfid_input_reader.py
