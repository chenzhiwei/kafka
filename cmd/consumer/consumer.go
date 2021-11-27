package main

import (
	"context"
	"fmt"
	"os"

	kafka "github.com/segmentio/kafka-go"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "myTopic",
	})
	defer r.Close()

	if err := r.SetOffset(0); err != nil {
		fmt.Printf("unable to set offset: %v\n", err)
		os.Exit(1)
	}

	ctx := context.Background()
	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			fmt.Printf("unable to read message: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, m.Key, m.Value)
	}
}
