### build the docker image for raspberry pi
FROM arm32v6/alpine:3.7
MAINTAINER @protenhan

COPY rpi-rfid-musicplayer-linux-arm-7 /rfid-musicplayer/

CMD /rfid-musicplayer/rpi-rfid-musicplayer-linux-arm-7
