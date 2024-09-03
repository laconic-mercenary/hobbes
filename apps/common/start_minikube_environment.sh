#!/bin/sh

# docker install test
docker system info

# start minikube
minikube start \
    --memory=8192 --cpus=4 \
    --kubernetes-version=v1.20.2 \
    --insecure-registry="localhost:5000" \
    --extra-config=apiserver.enable-admission-plugins="LimitRanger,NamespaceExists,NamespaceLifecycle,ResourceQuota,ServiceAccount,DefaultStorageClass,MutatingAdmissionWebhook"

## install the docker registry add-on
minikube addons enable registry

# create bridge to minikube's docker image registry
## the following will allow access from the local desktop to the 
## registry addon running on the minikube cluster
## it allows us to push docker images to the localhost:5000 endpoint
## and yet still be accessible to minikube

docker run --detach --rm -it --network=host alpine ash -c "apk add socat && socat TCP-LISTEN:5000,reuseaddr,fork TCP:$(minikube ip):5000"

