#!/bin/bash 

# prepare for cross building docker containers for ARM (Raspberry Pi)
docker run --rm --privileged multiarch/qemu-user-static:register --reset
for target_arch in aarch64 arm; do
  wget -N https://github.com/multiarch/qemu-user-static/releases/download/v3.0.0/x86_64_qemu-${target_arch}-static.tar.gz
  tar -xvf x86_64_qemu-${target_arch}-static.tar.gz
done