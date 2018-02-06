FROM golang:1.9-alpine3.7 as builder

ENV GOOS=linux
ENV GOARCH=arm
ENV GOARM=7

COPY src/go /go/src/musicplayer/

WORKDIR /go/src/musicplayer

RUN go build -v .

FROM hypriot/rpi-alpine
MAINTAINER @protenhan

COPY --from=builder /go/src/musicplayer/musicplayer /rfid-musicplayer/

RUN /rfid-musicplayer/musicplayer
