#!/bin/bash
set -e

image="protenhan/rpi-rfid-musicplayer"
docker tag rpi-rfid-musicplayer "$image:latest"
docker push "$image:latest"
docker tag rpi-rfid-musicplayer "$image:$TRAVIS_TAG-alpine-$ARCH"
docker push "$image:$TRAVIS_TAG-alpine-$ARCH"
curl --request POST https://hooks.microbadger.com/images/protenhan/rpi-rfid-musicplayer/liRpfIP6ir_P47tZi-rkKg64Shc\=
