#!/bin/sh

## a utility script that base64 encodes secret values 
## so that they can be placed in app secrets.yaml
## note: better to use VAULT

if [[ "$#" -ne 3 ]]; 
then
    echo "ERROR: incorrect number of arguments. "
    echo "USAGE: /bin/sh k8s_secretb64.sh <secret_name> <secret_key> <secret_value>"
    exit 1
fi

SECRETNAME=$1
SECRETKEY=$2
SECRETVALUE=$3

kubectl create secret generic ${SECRETNAME} \
    --from-literal ${SECRETKEY}=${SECRETVALUE} \
    -o yaml \
    --dry-run
