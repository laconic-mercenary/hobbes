#!/bin/sh

## builds and pushes to container registry on digital ocean

TAG="localhost:5000/ibkr-clientportal:latest"
docker build . --tag ${TAG}
docker push ${TAG}
