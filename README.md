Kafka
===

Get-Started with helm/kafka-ha on kubernetes

# Get-Started

```
go get github.com/chechiachang/kafka-playground
cd ${GOPATH}/github.com/chechiachang/kafka-playground
```

# Install

- [vagrant](vagrant)
- [docker-compose (confluent-kafka)](docker-compose)
- [helm (helm-kafka)](helm)

```
wget https://github.com/strimzi/strimzi-kafka-operator/releases/download/0.18.0/strimzi-cluster-operator-0.18.0.yaml
sed -i 's/namespace: .*/namespace: kafka/' strimzi-cluster-operator-0.18.0.yaml

kubectl create ns kafka
kubectl -n kafka apply -f strimzi-cluster-operator-0.18.0.yaml
```

# vagrant

# golang Client Example

```
vim client-go/main.go
```
