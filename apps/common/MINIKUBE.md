### Overview

The following is a simple operation manual for minikube.

### Usage

* Starting

```
minikube start --insecure-registry="localhost:5000"
```

* Stopping

```
minikube stop
```

* Delete Cluster

```
minikube delete
```

### About Docker Images

* Minikube runs in it's own 'docker space' - so docker commands that are executed on your terminal are not visible to minikube. 

* What this means is that even if you build images and push to your local registry - minikube is NOT able to pull from this registry.

#### Install Docker Registry add-on

* To use the local docker registry, an add-on must be installed in minikube. 
    * Follow these steps: https://minikube.sigs.k8s.io/docs/handbook/registry/#docker-on-macos
    * See also k8s/k8s_minikube_registry_bridge.sh

### Accessing a Service in minikube

* In order to access an service deployed in minikube, NodePort needs to be used.
    * More Info: https://minikube.sigs.k8s.io/docs/handbook/accessing/

* Each project should have a k8s_minikube_expose.sh script that explains more.