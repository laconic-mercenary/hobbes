### Overview

These are utilities and manifests used to deploy the science python functions to K8S.

### Prerequisites

* A K8S cluster with kubectl access

### Deploy

* Deploy the S&P500 Service
```
kubectl apply -f all-sp500.yaml
```

* Deploy the ALL Service
```
kubectl apply -f all-allmkt.yaml
```