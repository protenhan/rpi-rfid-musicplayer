#!/bin/bash
set -e

docker version
uname -a
echo "Updating Docker engine to have multi-stage builds"
sudo service docker stop
curl -fsSL https://get.docker.com/ | sudo sh
docker version

go get github.com/karalabe/xgo
xgo -go 1.9.2 --targets=linux/arm-* github.com/protenhan/rpi-rfid-musicplayer
ls -la rpi-rfid*
file rpi-rfid-musicplayer-linux-arm-7

if [ -d tmp ]; then
  docker rm build
  rm -rf tmp
fi

docker build -t rpi-rfid-musicplayer .
