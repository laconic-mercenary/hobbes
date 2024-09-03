#!/bin/sh

### enables you to reach the web app locally

echo "Kibana will be available at http://localhost:(the indicated port)"

## format is <local port>:<server port>
### NOTE: was having connectivity problems when specifying the local port
### but NOT specifying the port significantly reduced these problems
### so will not specify one for now
## see: https://github.com/kubernetes/kubernetes/issues/74551

kubectl port-forward -n elasticstack svc/kibana :5601
