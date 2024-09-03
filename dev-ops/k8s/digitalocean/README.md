### Overview

The guide on how to use Digital Ocean K8S as a service.

### Commands


#### DOCT
https://docs.digitalocean.com/reference/doctl/reference/


#### KUBECTL

Get Nodes

```
kubectl --context do-sfo3-hobbes-k8s-us-sf get nodes
```

#### Sample Deployment

https://github.com/digitalocean/doks-example

#### Reference

https://docs.digitalocean.com/products/kubernetes/

#### Docker Registry

* guide to using the container image registry:
https://docs.digitalocean.com/products/container-registry/quickstart/

* hobbes registry
https://cloud.digitalocean.com/registry?i=b534d8


* Login to Registry
``` 
doctl registry login 
```

* Tag Image 
```
docker tag <my-image> registry.digitalocean.com/hobbes/<my-image> 
```

* Push Image
```
docker push registry.digitalocean.com/hobbes/<my-image>
```

