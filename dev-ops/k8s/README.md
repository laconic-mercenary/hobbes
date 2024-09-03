### Overview

Here you will find all the systems and installation steps necessary to run the Hobbes services.

The services are deployed on a Kubernetes, OpenFaaS, Knative stack.

### Installation Steps

The order of installation is:

1. digitalocean 
    * Kubernetes (K8S)
    * Installation [here](./digialocean/README.md)

2. knative 
    * Powerful Kubernetes features 
    * Installation [here](./knative/README.md)

3. openfaas 
    * Functions As A Service
    * Installation [here](./openfaas/README.md)

4. K8S dashboard 
    * Dashboard for managing K8S apps
    * Installation [here](./dashboard/README.md)

5. elastic search + kibana + metric beat
    * Data and App logging repository
    * Installation [here](./elasticsearch/README.md)