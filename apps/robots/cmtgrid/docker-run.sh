#!/bin/bash

docker build . -t "robot:test"
docker run -rm --env-file=./env "robot:test"