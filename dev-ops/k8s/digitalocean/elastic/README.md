### OVERVIEW

The following will install the elastic stack on your Kubernetes cluster. All containers deployed to this Kubernetes will then have their logs visible on Kibana. 

### NOTES

* Elastic Search is in SINGLE node mode in order to save resources
* Kibana will be accessed via ``` kubectl port-forward ``` rather than a public load balancer

### PREREQUISITES

* A Kuberentes cluster
* A ``` kubectl ``` that can connect to the Kubernetes cluster

### INSTALLATION

1. Create common namespace
* Deploy
```
kubectl apply -f namespace.yaml
```

2. ELASTICSEARCH
* Deploy
```
kubectl apply -f elasticsearch/all.yaml
```

3. KIBANA
* Deploy
```
kubectl apply -f kibana/all.yaml
```

4. LOGSTASH
* Deploy
```
kubectl apply -f logstash/all.yaml
```

5. FILEBEAT
* Deploy
```
kubectl apply -f filebeat/all.yaml
```

6. Access Kibana
* Port Forward
```
sh kibana-portfwd.sh
```

7. Add filebeat INDEX to Kibana
    * Navigate /app/management/kibana/indexPatterns/create
    * In Step 1, Create the following index ``` filebeat* ```
    * In Step 2, Select ``` @timestamp ```

8. Modify the INDEX 
    * Navigate /app/management/data/index_lifecycle_management/policies/edit/filebeat
    * Expand ``` Advanced Settings ```
    * Change Maximum Age from 30 to 3 days
    * Change Maximum index size from 50 to 1GB

9. ELASTALERT
TODO


### REFERENCE

* https://coralogix.com/blog/elasticsearch-logstash-kibana-on-kubernetes/

