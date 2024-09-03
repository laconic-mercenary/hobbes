#!/bin/sh

### when using the registry addon with miniukbe, make sure to keep this running so you can reach the registry locally

docker run --rm -it --network=host alpine ash -c "apk add socat && socat TCP-LISTEN:5000,reuseaddr,fork TCP:$(minikube ip):5000"