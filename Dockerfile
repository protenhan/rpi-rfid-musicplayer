FROM golang:1.9.2-alpine3.7 as builder

# install git to enable go get
RUN apk update
RUN apk add git

ENV GOPATH=$GOPATH:/musicplayer

COPY ./ /musicplayer/
WORKDIR /musicplayer/

# resolve dependencies and build the binaries
RUN go get github.com/karalabe/xgo
RUN xgo -go 1.9.2 --targets=linux/arm-7 build -v github.com/protenhan/rpi-rfid-musicplayer

### build the docker image for raspberry pi
FROM arm32v6/alpine:3.7
MAINTAINER @protenhan

COPY --from=builder /go/src/musicplayer/musicplayer /rfid-musicplayer/

RUN /rfid-musicplayer/musicplayer
