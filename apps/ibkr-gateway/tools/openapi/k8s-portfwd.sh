#!/bin/sh
kubectl port-forward -n default svc/swagger-ui 33331:8080