package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := os.Getenv("KAFKA_TOPIC")
	partition, err := strconv.Atoi(os.Getenv("KAFKA_PARTITION"))
	if err != nil {
		log.Println("Failed to parse env KAFKA_PARTITION")
		partition = 0
	}
	kafkaURL := os.Getenv("KAFKA_URL")
	log.Printf("topic: %s partition: %v kafkaURL: %s", topic, partition, kafkaURL)

	producerConn, err := kafka.DialLeader(context.Background(), "tcp", kafkaURL, topic, partition)
	if err != nil {
		panic(err)
	}
	defer producerConn.Close()

	//producerConn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	go func() {
		for {
			producerConn.WriteMessages(
				kafka.Message{
					Value: []byte(strconv.Itoa(time.Now().Second())),
				},
			)
			time.Sleep(1 * time.Second)
		}
	}()

	// make a new reader that consumes from topic-A, partition 0, at offset 42
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{kafkaURL},
		Topic:     topic,
		Partition: 0,
		MinBytes:  10e2, // 1KB
		MaxBytes:  10e3, // 10KB
	})
	defer r.Close()
	//r.SetOffset(42)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println(err)
			break
		}
		log.Println("%v message at offset %d: %s = %s\n", time.Now(), m.Offset, string(m.Key), string(m.Value))
	}

}
