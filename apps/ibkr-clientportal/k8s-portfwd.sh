#!/bin/sh

### enables you to reach the web app locally

## format is <local port>:<server port>
kubectl port-forward -n hobbes svc/ibkr-clientportal 31004:5001