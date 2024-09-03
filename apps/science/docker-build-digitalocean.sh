#!/bin/sh

TAG="registry.digitalocean.com/hobbes/science:latest"

docker build . -f=Dockerfile -t=${TAG}

docker push ${TAG}