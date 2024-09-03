#!/bin/sh
export OPENFAAS_URL=http://127.0.0.1:8080

OFP=$(kubectl get secret -n openfaas basic-auth -o jsonpath="{.data.basic-auth-password}" | base64 --decode; echo)

faas-cli login --username admin --password ${OFP} --gateway "http://127.0.0.1:8080"

faas-cli deploy --gateway "http://127.0.0.1:8080" -f ibkr-gateway.yml