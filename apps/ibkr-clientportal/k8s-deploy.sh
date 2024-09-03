#!/bin/sh

## this will update k8s objects, if there are any changes

function apply() {
    kubectl apply -f $1
}

apply k8s/namespace.yaml
apply k8s/configmap.yaml
apply k8s/secrets.yaml
apply k8s/service.yaml
apply k8s/deployment.yaml
