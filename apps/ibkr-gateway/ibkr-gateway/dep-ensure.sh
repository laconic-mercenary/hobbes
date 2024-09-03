#!/bin/sh

## imports new modules from github.com
## simply import them in the go source 
# and then run this script

rm -rf ./vendor
docker build . -f=Dockerfile-Build --tag=localhost:5000/ibkr-gateway:build
docker run --rm -v $(PWD):/go/src/ibkr-gateway localhost:5000/ibkr-gateway:build