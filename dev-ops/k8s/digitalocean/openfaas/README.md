### Overview

OpenFaaS installation on Digital Ocean K8S.

### Prerequisites

1. Install HELM 

### Installation

1. Create namespaces
```
kubectl apply -f namespace.yaml
```

2. Pull HELM OpenFaaS 
```
helm repo add openfaas https://openfaas.github.io/faas-netes/
```

3. Install OpenFaaS via HELM
```
helm --kubeconfig=~/.kube/config repo update && helm upgrade openfaas --install openfaas/openfaas \
    --namespace openfaas --set functionNamespace=openfaas-fn --set generateBasicAuth=true
```

4. Install FaaS CLI
    * faas-cli: https://docs.openfaas.com/cli/install/#linux-or-macos

5. Port Foward to access the OpenFaaS (do this in separate, dedicated terminal window)
```
kubectl port-forward -n openfaas svc/gateway 8080:8080
```

6. Login to OpenFaaS
```
sh openfaas-login.sh
```

