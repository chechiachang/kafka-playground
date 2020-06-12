
[Benchmark Command](https://gist.github.com/jkreps/c7ddb4041ef62a900e6c)

# 1 master 3 n1-standard-1

```
TOPIC=test
ZOOKEEPER=its-cdh-kafka-tw-01-01:2181
/usr/bin/kafka-topics --zookeeper ${ZOOKEEPER} --create --topic ${TOPIC} --partitions 9 --replication-factor 3
```

Producer
```
BOOTSTRAP_SERVERS=its-cdh-kafka-tw-01-01:9092,its-cdh-kafka-tw-01-02:9092,its-cdh-kafka-tw-01-03:9092
THROUGHPUT=4000
NUM_RECORDS=$((${THROUGHPUT} * 100))
RECORD_SIZE=1024
BATCH_SIZE=1024
/usr/bin/kafka-producer-perf-test \
  --num-records ${NUM_RECORDS} \
  --record-size ${RECORD_SIZE} \
  --topic ${TOPIC} \
  --throughput ${THROUGHPUT} \
  --producer-props bootstrap.servers=${BOOTSTRAP_SERVERS} \
    max.in.flight.requests.per.connection=5 \
    batch.size=${BATCH_SIZE} \
    compression.type=none

400000 records sent, 3997.162015 records/sec (3.90 MB/sec), 194.64 ms avg latency, 2600.00 ms max latency, 6 ms 50th, 1610 ms 95th, 2409 ms 99th, 2538 ms 99.9th.

THROUGHPUT=5000
NUM_RECORDS=$((${THROUGHPUT} * 100))
400000 records sent, 4998.500450 records/sec (4.88 MB/sec), 177.38 ms avg latency, 1911.00 ms max latency, 9 ms 50th, 1475 ms 95th, 1732 ms 99th, 1835 ms 99.9th.

THROUGHPUT=6000
NUM_RECORDS=$((${THROUGHPUT} * 100))
400000 records sent, 5995.922773 records/sec (5.86 MB/sec), 1456.71 ms avg latency, 4032.00 ms max latency, 1322 ms 50th, 3622 ms 95th, 3829 ms 99th, 3958 ms 99.9th.

RECORD_SIZE=8192
BATCH_SIZE=8192
THROUGHPUT=5000
NUM_RECORDS=$((${THROUGHPUT} * 100))
500000 records sent, 4993.857555 records/sec (39.01 MB/sec), 543.23 ms avg latency, 2488.00 ms max latency, 128 ms 50th, 2028 ms 95th, 2228 ms 99th, 2448 ms 99.9th.

RECORD_SIZE=16384
BATCH_SIZE=16384
THROUGHPUT=4000
NUM_RECORDS=$((${THROUGHPUT} * 100))
400000 records sent, 3934.258540 records/sec (61.47 MB/sec), 494.55 ms avg latency, 8879.00 ms max latency, 227 ms 50th, 1317 ms 95th, 4038 ms 99th, 8711 ms 99.9th.
```

# Consumer

```
BATCH_SIZE=100
/usr/bin/kafka-consumer-perf-test \
  --bootstrap-server ${BOOTSTRAP_SERVERS} \
  --messages 50000 \
  --fetch-size ${BATCH_SIZE} \
  --topic ${TOPIC} \
  --threads 1 \
  --timeout 60000

start.time, end.time, data.consumed.in.MB, MB.sec, data.consumed.in.nMsg, nMsg.sec, rebalance.time.ms, fetch.time.ms, fetch.MB.sec, fetch.nMsg.sec

BATCH_SIZE=8196
```

