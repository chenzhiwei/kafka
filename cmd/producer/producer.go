package main

import (
	"context"
	"fmt"
	"os"

	"github.com/segmentio/kafka-go"
)

func main() {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
	})
	defer w.Close()

	ctx := context.Background()
	if err := w.WriteMessages(ctx, kafka.Message{
		Topic: "myTopic",
		Key:   []byte("kafka"),
		Value: []byte("kafka"),
	}); err != nil {
		fmt.Printf("unable to write message: %v", err)
		os.Exit(1)
	}
}
