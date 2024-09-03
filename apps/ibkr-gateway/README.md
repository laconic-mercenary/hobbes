### Overview

The function was generated using faas-cli, using the following steps:

1. ``` faas-cli template store pull golang-http ```
2. ``` faas-cli new --lang golang-http ibkr-gateway ```

Note: the downloaded template was put in source control for stability reasons.
* see template/golang-http directory

### Prerequisites

* OpenFaaS
    * https://docs.openfaas.com/cli/install/
* Kubernetes
    * Install K8S using modern versions of Docker
    * Or use https://minikube.sigs.k8s.io/docs/start/
* A Local Docker Registry 
    * running on port 5000
* Go Dep 
    * Install from here https://github.com/golang/dep
    * Note: using 'dep' instead of go modules because of the issues encountered getting it to work with openfaas.


### Deploy

This function is designed to be deployed to Kubernetes. 

1. Fetch Template (ALREADY DONE - see template/golang-http directory)

OpenFaaS relies on templates - the following will pull the golang template w/ http support.
Take note of the golang version in the Dockerfile within the template.

``` faas-cli template store pull golang-http ```

2. Fetch the Vendor directory

Run the script found in the same directory that handler.go is found.

``` /bin/sh run-dep-ensure.sh ```

3. Build and Push Image

``` /bin/sh faas-build-local.sh ```

4. K8S Port Forwarding

Enables us to interact with the OpenFaaS API.

``` /bin/sh k8s-portfwd.sh ``` 

5. Deploy to K8S

``` /bin/sh k8s-deploy.sh ```


### Usage

#### App Logs in K8S

```
PODID=$(kubectl get pods -n openfaas-fn -o jsonpath='{range .items[*]} {.metadata.name}' | grep ibkr-gateway)
kubectl logs ${PODID} -n openfaas-fn
```

#### Local Testing

As this is a function, there is no main method to launch with ``` go run ```. Instead use ``` go test ``` and ``` go fmt ```. Both of those utilities come with a golang installation: https://golang.org/doc/install

* Check Syntax

``` go test handler.go ```

* Format Source Code

``` go fmt handler.go ``` 


#### Adding New Modules

Although the go modules is the modern way to manage dependencies in golang, there was great difficulty faced in trying to get it to work. As such, the older go deps approach was used. 

To add a new library, use dep ensure -add and the name of the library you want. For example, to add ``` github.com/cnf/structhash ``` for use in our function, run:

``` 
dep ensure -add github.com/cnf/structhash
```

Now reference the package using the typical an import statement

```
import "github.com/cnf/structhash"
```

You an also import it in your code and execute ``` dep ensure ``` but should remove the vendor directory first.

If that doesn't work, update the Gopkg.toml manually, then run the "run_dep_ensure.sh" script.

Note: the openfaas dependency ``` "github.com/openfaas/templates-sdk/go-http" ``` must NOT exist in the Gopkg.toml - or it will overwrite the one found in 'vendor' in the openfaas go-lang template.
