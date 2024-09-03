### Overview

This is the clientportal web application that fronts the IBKR account services. 

Important: I do NOT own this application.

### IBKR Capabilities

The clientportal exposes an API. For more info:
https://interactivebrokers.github.io/cpwebapi/

For API doc:
https://interactivebrokers.com/api/doc.html

### Deploying to K8S

The following steps will deploy the clientportal to Kubernetes.

1. Verify Configuration
    * root/conf.yaml
        * Check the ips.allow group to make sure access is appropriately restricted.

2. Build and push the local docker image
    * Run docker_build_local.sh

3. Deploy to K8S
    * Run k8s/k8s_deploy.sh

### (Minikube) Accessing the Service

Accessing the service through the NodePort on minikube requires another step.

* Run the following (blocking)
```
/bin/sh k8s/k8s_minikube_expose.sh
```

### Usage

Open a web browser to the target address and login with your IBKR credentials.

This allows other apps 

### TLS

There were many problems getting localhost HTTPS to work on the clients - as a result, "listenSsl" is set to false in the conf.yaml. To enable this later, set this to true and see the ``` refresh_tls_certificate.sh ``` script.

### Helpful K8S Commands

* Get all clientportal pods

```
kubectl get pods -n hobbes-ibkr
```

* Get stdout logs a clientportal pod

```
kubectl logs clientportal-<pod id> -n hobbes-ibkr
```

* View additional logs in clientportal pod

```
kubectl exec -it clientportal-<pod id> -n hobbes-ibkr -- /bin/sh
cd logs
ls -la
```

* Delete Pod

```
kubectl delete pod clientportal-7d7b777c58-dpstj -n hobbes-ibkr
```

* Delete Deployment

```
kubectl delete deployment clientportal -n hobbes-ibkr
```