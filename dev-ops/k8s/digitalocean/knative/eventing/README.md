### OVERVIEW

This installs the following:

1) Apache Kafka (Strimzi)
2) Knative Apache support

This enables you to use Apache Kafka as the messaging implementation of Knative Eventing.

### INSTALLATION

1. Create NS
```
kubectl create namespace kafka
```

2. Create strimzi operator
```
kubectl create -f kafka/strimzi-cluster-operator-0.24.0.yaml -n kafka
```

3. Create the hobbes kafka cluster
```
kubectl apply -f kafka/kafka-cluster.yaml -n kafka
```

4. Monitor the cluster status
```
kubectl wait kafka/hobbes-kafka-cluster --for=condition=Ready --timeout=300s -n kafka
```

5. Create the KN Kafka Broker
```
kubectl apply -f kn-kafka/default-config.yaml
kubectl apply -f kn-kafka/default-broker.yaml
kubectl create -f kn-kafka/kafka-source-schema.yaml
kubect get brokers -n hobbes
kubectl get pods --namespace knative-sources
```

6. Create the KN Kafka Channels
```
kubectl create -f kn-kafka-channels/kafka-channel-schema.yaml
kubectl create -f kn-kafka-channels/default-ch-webhook.yaml
```

### REFERENCES

* Details can be found here (highly suggested reading for additional configuration)
    https://strimzi.io/docs/operators/latest/using.html

* https://strimzi.io/quickstarts/

* https://knative.dev/docs/eventing/samples/kafka/channel/

* https://developingfordata.com/tag/eventing/
