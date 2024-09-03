### OVERVIEW

How to deploy ElastAlert on Kubernetes.

### PREREQUISITES

### RULES

Rules in ElastAlert are configured via yaml files. In this case, they are supplied via a Kubernetes ConfigMap. 

[See ConfigMap 'ea-rules' here](./all.yaml)

Feel free to add as many rules as you want. These files will mount to the container directory that ElastAlert expects to find rules in.

Rules use the Lucene language for querying - and is easily tested on Kibana.

### INSTALL

* Deploy
```
kubectl apply -f all.yaml
```

### MAKING CHANGES
If any changes need to be made to all.yaml, take the following steps:

1. Apply the new changes
```
kubectl apply -f all.yaml
```

2. Delete the old pod
```
## get the pod ID first
kubectl get pods -n elasticstack | grep elastalert
## delete
kubectl delete pod (POD ID) -n elasticstack
```

3. Verify changes are ok
```
## get the pod ID first
kubectl get pods -n elasticstack | grep elastalert
## view the logs of the pod
kubectl logs (POD ID) -n elasticstack
```

### TROUBLESHOOTING

* Getting running pods of elastalert
```
kubectl get pods -n elasticstack | grep elastalert
```

* Viewing logs of a elastalert (this is useful for debugging rules)
```
## get the pod ID first
kubectl get pods -n elasticstack | grep elastalert
## view the logs of the pod
kubectl logs (POD ID) -n elasticstack
```