
# Guide

https://strimzi.io/quickstarts/

# Install operator

Download operator manefest
```
REPO=git@github.com:strimzi/strimzi-kafka-operator.git
VERSION=0.18.0
kpt pkg get ${REPO}/install/cluster-operator@${VERSION} cluster-operator
```

Set namespace to kafka
```
sed -i '' 's/namespace: .*/namespace: kafka/' install/cluster-operator/*RoleBinding*.yaml

vim install/cluster-operator/050-Deployment-strimzi-cluster-operator.yaml
env:
- name: STRIMZI_NAMESPACE
  value: my-kafka-project
```

Apply
```
kubectl create namespace kafka

kubectl apply -f cluster-operator/ -n kafka
```

Update
```
VERSION=0.18.0
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
kubectl -n kafka run kafka-producer -ti --image=strimzi/kafka:0.18.0-kafka-2.5.0 --rm=true --restart=Never -- bin/kafka-console-producer.sh --broker-list my-cluster-kafka-bootstrap:9092 --topic my-topic

kubectl -n kafka run kafka-consumer -ti --image=strimzi/kafka:0.18.0-kafka-2.5.0 --rm=true --restart=Never -- bin/kafka-console-consumer.sh --bootstrap-server my-cluster-kafka-bootstrap:9092 --topic my-topic --from-beginning
```
