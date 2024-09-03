#!/bin/bash

### this will deploy all middleware services to the 
## k8s cluster
## prerequisites are listed below:
kubectl get nodes
helm
istioctl
faas-cli
file ~/.kube/config

## begin
kubectl create namespace hobbes

## serving
kubectl create namespace istio-operator
istioctl install --verify

kubectl create -f knative/serving/serving-crds.yaml
kubectl create -f knative/serving/serving-core.yaml

kubectl rollout status Deployment/controller -n knative-serving
kubectl rollout status Deployment/autoscaler -n knative-serving
kubectl rollout status Deployment/activator -n knative-serving

## eventing
kubectl create -f knative/eventing/eventing-crds.yaml
kubectl create -f knative/eventing/eventing-core.yaml

kubectl rollout status Deployment/eventing-controller -n knative-eventing
kubectl rollout status Deployment/eventing-webhook -n knative-eventing

kubectl create namespace kafka
kubectl create -f knative/eventing/kafka/strimzi-cluster-operator-0.24.0.yaml -n kafka
kubectl rollout status Deployment/strimzi-cluster-operator -n kafka

kubectl create -f knative/eventing/kafka/kafka-cluster.yaml -n kafka
kubectl wait kafka/hobbes-kafka-cluster --for=condition=Ready --timeout=300s -n kafka

kubectl create -f knative/eventing/kn-kafka/default-config.yaml
kubectl create -f knative/eventing/kn-kafka/default-broker.yaml
kubectl create -f knative/eventing/kn-kafka/kafka-source-schema.yaml

kubectl get brokers -n hobbes
kubectl rollout status Deployment/kafka-controller-manager -n knative-sources

kubectl apply -f knative/eventing/kn-kafka-channels/default-ch-webhook.yaml
kubectl create -f knative/eventing/kn-kafka-channels/kafka-channel-schema.yaml

## openfaas
helm repo add openfaas https://openfaas.github.io/faas-netes/
kubectl create namespace openfaas
kubectl create namespace openfaas-fn

helm --kubeconfig=~/.kube/config repo update && helm upgrade openfaas --install openfaas/openfaas \
    --namespace openfaas --set functionNamespace=openfaas-fn --set generateBasicAuth=true
kubectl -n openfaas get deployments -l "release=openfaas, app=openfaas"
kubectl rollout status Deployment/gateway -n openfaas

echo ">>> NOTE:"
echo "kubectl port-forward -n openfaas svc/gateway 8080:8080"
echo "/bin/sh openfaas/openfaas-login.sh"

## elasticstack
kubectl create namespace elasticstack

kubectl create -f elastic/elasticsearch/all.yaml
kubectl rollout status Deployment/elasticsearch -n elasticstack

kubectl create -f elastic/logstash/all.yaml
kubectl rollout status Deployment/logstash -n elasticstack

kubectl create -f elastic/filebeat/all.yaml
kubectl rollout status DaemonSet/filebeat -n kube-system

kubectl create -f elastic/elastalert/all.yaml
kubectl rollout status Deployment/elastalert -n elasticstack

kubectl exec `kubectl get pod --selector=app=elasticsearch --output=jsonpath={.items..metadata.name} -n elasticstack` \ 
    -n elasticstack -- curl -XPUT http://localhost:9200/elastalert_status

kubectl create -f elastic/kibana/all.yaml
kubectl rollout status Deployment/kibana -n elasticstack

## k8s dashboard
kubectl create -f dashboard/all.yaml
kubectl rollout status Deployment/kubernetes-dashboard -n kubernetes-dashboard
echo ">>> NOTE:"
echo "/bin/sh dashboard/k8s_print_token.sh"

## done
kubectl get pods --all-namespaces

echo "FINISHED"