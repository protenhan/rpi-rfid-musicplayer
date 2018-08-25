FROM arm32v6/alpine:3.8
LABEL maintainer="@protenhan"

COPY qemu-arm-static /usr/bin/
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

# Setup for the project
COPY src/python /rfid-musicplayer/
RUN pip3 install -r /rfid-musicplayer/requirements.txt

CMD python3 /rfid-musicplayer/rfid_input_reader.py
