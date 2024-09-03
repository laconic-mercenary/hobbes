#!/bin/sh

kubectl --version
helm --version

kubectl apply -f namespace.yaml

helm repo add openfaas https://openfaas.github.io/faas-netes/

helm --kubeconfig=~/.kube/config repo update && helm upgrade openfaas --install openfaas/openfaas \
    --namespace openfaas --set functionNamespace=openfaas-fn --set generateBasicAuth=true

