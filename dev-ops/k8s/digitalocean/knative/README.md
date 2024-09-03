### Overview

Steps to install knative serving and eventing on your kubernetes cluster.

### Installation

1. Create Core Serving
```
kubectl apply -f serving/serving-crds.yaml
kubectl apply -f serving/serving-core.yaml
kubectl apply -f serving/istio-minimal-operator.yaml
kubect get pods -n knative-serving
```

2. Create Core Eventing
```
kubectl apply -f eventing-crds.yaml
kubectl apply -f eventing-core.yaml
kubect get pods -n knative-eventing
```

3. Deploy Kafka and Eventing
[here](./eventing/README.md)

4. Deploy Broker
