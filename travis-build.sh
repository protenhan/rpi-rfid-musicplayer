#!/bin/bash
set -e

docker version
uname -a
echo "Updating Docker engine to have multi-stage builds"
sudo service docker stop
curl -fsSL https://get.docker.com/ | sudo sh

if [ -d tmp ]; then
  docker rm build
  rm -rf tmp
fi

docker build -t rpi-rfid-musicplayer .
