#!/bin/sh

TAG="localhost:5000/swagger-ui:latest"
docker build . --tag ${TAG}
docker push ${TAG}
