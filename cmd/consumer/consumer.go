package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	kafka "github.com/segmentio/kafka-go"

	v1 "github.com/chenzhiwei/kafka/api/v1"
)

var (
	broker string
	topic  string
	group  string
	offset int64
)

func main() {
	config := kafka.ReaderConfig{
		Brokers: strings.Split(broker, ","),
	}

	if group == "" {
		config.Topic = topic
	} else {
		config.GroupID = group
		config.GroupTopics = strings.Split(topic, ",")
	}

	r := kafka.NewReader(config)
	defer r.Close()

	if offset >= 0 {
		if err := r.SetOffset(offset); err != nil {
			fmt.Printf("unable to set offset: %v\n", err)
			os.Exit(1)
		}
	}

	ctx := context.Background()
	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			fmt.Printf("unable to read message: %v\n", err)
			os.Exit(1)
		}

		var headers []v1.Header
		for _, header := range m.Headers {
			headers = append(headers, v1.Header{Key: header.Key, Value: string(header.Value)})
		}
		fmt.Printf("Message at %s@%d[%d], Key=%s, Value=%s, Headers=%v\n", m.Topic, m.Partition, m.Offset, m.Key, m.Value, headers)
	}
}

func init() {
	flag.StringVar(&broker, "broker", "localhost:9092", "kafka brokers seperated by comma")
	flag.StringVar(&topic, "topic", "", "kafka brokers seperated by comma")
	flag.StringVar(&group, "group-id", "", "kafka group id")
	flag.Int64Var(&offset, "offset", -1, "kafka topic offset")
	flag.Parse()

	if topic == "" {
		panic("topic must not be empty")
	}

	if group == "" && strings.Contains(topic, ",") {
		panic("subscribe to multiple topics requires setting a group id")
	}
}
