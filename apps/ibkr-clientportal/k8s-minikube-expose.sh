#!/bin/sh

## the following will create a tunnel to the nodeport for the clientportal
## assuming that it's using a nodeport and not a clusterIPs
## typically this is used if wanting to access the clientportal directly 
## from the desktop - for example, a webbrowser

minikube service --url ibkr-clientportal -n hobbes