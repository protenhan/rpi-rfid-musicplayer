#!/bin/bash
set -e

image="protenhan/rpi-rfid-musicplayer"
# push tagged images if the commit was tagged
if [ $# -eq 1 ]; then
  docker tag rpi-rfid-musicplayer "$image:$TRAVIS_TAG-alpine-$ARCH"
  docker push "$image:$TRAVIS_TAG-alpine-$ARCH"
fi
 
# always push a latest image
docker tag rpi-rfid-musicplayer "$image:latest"
docker push "$image:latest"
  
