#!/bin/sh

## prints the token that allows you to access the 
## K8S dashboard 

## NOTE: make sure to create the accounts in account.yaml 
## first

kubectl -n kubernetes-dashboard get secret \
    $(kubectl -n kubernetes-dashboard get sa/admin-user -o jsonpath="{.secrets[0].name}") \
    -o go-template="{{.data.token | base64decode}}"

