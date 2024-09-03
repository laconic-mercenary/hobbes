#!/bin/bash

set -e

## faas-cli is required!
faas-cli version

export OPENFAAS_URL="http://127.0.0.1:8080"

OFP=$(kubectl get secret -n openfaas basic-auth -o jsonpath="{.data.basic-auth-password}" | base64 --decode; echo)

faas-cli login --username admin --password ${OFP} --gateway "${OPENFAAS_URL}"

echo "> Username: admin"
echo "> Password: ${OFP}"