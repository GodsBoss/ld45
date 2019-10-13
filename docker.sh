#!/bin/bash

if [ "$1" == "prepare" ]
then
  docker build -t godsboss/gopherjs .
fi

if [ "$1" == "run" ]
then
  docker run \
    --rm \
    -it \
    -u 1000:1000 \
    --mount type=bind,src=${PWD},dst=/go/src/github.com/GodsBoss/ld45 \
    -p 8080:8080 \
    --workdir /go/src/github.com/GodsBoss/ld45 \
    godsboss/gopherjs \
    gopherjs serve -v github.com/GodsBoss/ld45/main/ld45
fi
