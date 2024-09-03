#!/bin/sh

TAG="registry.digitalocean.com/hobbes/ibkr-events:latest"

docker build . -f=Dockerfile -t=${TAG}

## attempt to push to local registry - it's ok if the registry is not running
## and this fails (if that was the intention)
docker push ${TAG}