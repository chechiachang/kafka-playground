package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "ticker"
	partition := 0
	kafkaURL := "localhost:9092"

	producerConn, _ := kafka.DialLeader(context.Background(), "tcp", kafkaURL, topic, partition)
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
			break
		}
		fmt.Printf("%v message at offset %d: %s = %s\n", time.Now(), m.Offset, string(m.Key), string(m.Value))
	}

}
