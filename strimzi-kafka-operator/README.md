
# Guide

https://strimzi.io/quickstarts/

# Install operator

Download operator manefest
```
REPO=git@github.com:strimzi/strimzi-kafka-operator.git
VERSION=0.17.0
kpt pkg get ${REPO}/install/cluster-operator@${VERSION} ${VERSION}
```

Set namespace to kafka
```
sed -i '' 's/namespace: .*/namespace: kafka/' ${VERSION}/*RoleBinding*.yaml

vim ${VERSION}/050-Deployment-strimzi-cluster-operator.yaml
env:
- name: STRIMZI_NAMESPACE
  value: my-kafka-project
```

Apply
```
kubectl create namespace kafka

kubectl apply -f 0.17.0 -n kafka
```

Update
```
VERSION=0.17.0
kpt pkg update cluster-operator@${VERSION}
```

---
Use operator

# Deploy kafka

```
kubectl apply -f kafka-persistent-single.yaml -n kafka
```

# Use kafka

```
kubectl -n kafka run kafka-producer -ti --image=strimzi/kafka:0.17.0-kafka-2.4.0 --rm=true --restart=Never -- bin/kafka-console-producer.sh --broker-list kafka-1-kafka-bootstrap:9092 --topic my-topic

kubectl -n kafka run kafka-consumer -ti --image=strimzi/kafka:0.17.0-kafka-2.4.0 --rm=true --restart=Never -- bin/kafka-console-consumer.sh --bootstrap-server kafka-1-kafka-bootstrap:9092 --topic my-topic --from-beginning
```
