### Overview

This deploys an administrative dashboard to manage Kubernetes.

For more info: https://kubernetes.io/docs/tasks/access-application-cluster/web-ui-dashboard/

### Installation

1. Run the following

```
kubectl apply -f all.yaml
```

2. Print the login TOKEN

```
/bin/sh k8s_print_token.sh
```

### Accessing

1. Open a proxy

```
kubectl proxy
```

2. Then navigate to the following

``` http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#/login ```


3. Use the token printed in k8s_print_token.sh to login
NOTE: do not include the '%' at the end.