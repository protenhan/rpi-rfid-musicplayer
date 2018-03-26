FROM golang:1.10.0-alpine3.7 as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=arm
ENV GOARM=7
ENV GOHOSTARCH=amd64

# install git to enable go get
RUN apk update
RUN apk add git

COPY ./ /musicplayer/
WORKDIR /musicplayer/

# resolve dependencies and build the binaries
RUN go build .

### build the docker image for raspberry pi
FROM arm32v6/alpine:3.7
MAINTAINER @protenhan

COPY --from=builder /musicplayer/musicplayer /rfid-musicplayer/

CMD /rfid-musicplayer/musicplayer
