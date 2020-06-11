build:
	skaffold build

run:
	skaffold run --tail

delete:
	skaffold delete

consumer:
	kubectl -n kafka run kafka-consumer -ti --image=strimzi/kafka:0.17.0-kafka-2.4.0 --rm=true --restart=Never -- bin/kafka-console-consumer.sh --bootstrap-server kafka-3-kafka-bootstrap:9092 --topic ticker --partition 0
