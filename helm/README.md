Kafka
===

Get-Started with helm/kafka-ha on kubernetes

# Usage

Read install.sh before usage.

```
chmod +x ./install.sh
./install.sh
```

# golang Client Example

```
vim main.go
```

# zookeeper

https://zookeeper.apache.org/doc/r3.3.3/zookeeperAdmin.html

```
kubectl exec -it kafka-zookeeper-0 --container kafka-broker bash
/usr/bin/zkCli.sh -server localhost:2181

# List root Nodes
ls /

# List /brokers/ids
ls /brokers/ids

```

# Kafka

```
kubectl exec -it kafka-0 --container kafka-broker bash

# List topics
/usr/bin/kafka-topics --list --zookeeper kafka-1-zookeeper

# List broker
/usr/bin/kafka-console-producer \
--broker-list localhost:9092\
 --topic topic-name

# This will create a new console-consumer and start consuming message to stdout
/usr/bin/kafka-console-consumer \
--bootstrap-server localhost:9092 \
--topic engine_topic_soundwave_USD \
--timeout 0 \
--from-beginning

# Use consumer to check trading-engine topics
/usr/bin/kafka-console-consumer \
--bootstrap-server localhost:9092 \
--topic engine_topic_JOGGP \
--group test

# Check consumer group
kafka-consumer-groups \
--bootstrap-server localhost:9092 \
--group cancel_result_topic_DAFGN_default_consumer \
--describe

kafka-consumer-groups \
--bootstrap-server localhost:9092 \
--group engine_topic_DAFGN_default_consumer \
--describe
```

# Config

https://kafka.apache.org/documentation/#topicconfigs
```
/usr/bin/kafka-configs --zookeeper kafka-zookeeper:2181 --describe max.message.bytes --entity-type topics

TOKEN=JZXTD
TOPIC=match_result_topic_${TOKEN}
TOPIC=engine_topic_${TOKEN}
TOPIC=cancel_result_topic_${TOKEN}
TOPIC=socketio_topic_${TOKEN}

/usr/bin/kafka-configs \
  --zookeeper kafka-zookeeper:2181 \
  --entity-type topics \
  --alter \
  --entity-name ${TOPIC} \
  --add-config max.message.bytes=16000000

```

# Performance Test

Producer
```
/usr/bin/kafka-producer-perf-test --num-records 100 --record-size 100 --topic performance-test --throughput 100 --producer-props bootstrap.servers=kafka:9092 max.in.flight.requests.per.connection=5 batch.size=100 compression.type=none

100 records sent, 99.108028 records/sec (0.01 MB/sec), 26.09 ms avg latency, 334.00 ms max latency, 5 ms 50th, 70 ms 95th, 334 ms 99th, 334 ms 99.9th.
```

Consumer
```
/usr/bin/kafka-consumer-perf-test --messages 100 --broker-list=kafka:9092 --topic performance-test --group performance-test --num-fetch-threads 1
```

# Operations

helm upgrade -f values-staging.yaml kafka incubator/kafka
