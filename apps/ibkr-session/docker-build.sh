#!/bin/sh

docker build ./refresh -f=Dockerfile -t=localhost:5000/ibkr-session:latest

## attempt to push to local registry - it's ok if the registry is not running
## and this fails (if that was the intention)
docker push localhost:5000/ibkr-session:latest