#!/bin/sh

kubectl port-forward -n openfaas svc/gateway 8080:8080