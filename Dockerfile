FROM golang:1.9.3-alpine3.7 as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=arm
ENV GOARM=7

# install git to enable go get
RUN apk update
RUN apk add git

COPY ./ /musicplayer/
WORKDIR /musicplayer/src/go

# resolve dependencies and build the binaries
RUN go get github.com/gvalkov/golang-evdev
RUN go build .

### build the docker image for raspberry pi
FROM arm32v6/alpine:3.7
MAINTAINER @protenhan

COPY --from=builder /go/src/musicplayer/musicplayer /rfid-musicplayer/

RUN /rfid-musicplayer/musicplayer
