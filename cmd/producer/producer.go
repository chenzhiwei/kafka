package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	kafka "github.com/segmentio/kafka-go"
	"gopkg.in/yaml.v3"

	v1 "github.com/chenzhiwei/kafka/api/v1"
	"github.com/chenzhiwei/kafka/utils/message"
)

var file string

func main() {
	bytes, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	data := &v1.Data{}
	if err := yaml.Unmarshal(bytes, data); err != nil {
		panic(err)
	}

	brokers := strings.Split(data.Broker, ",")
	messages, err := message.ToKafkaMessages(data)
	if err != nil {
		panic(err)
	}

	w := &kafka.Writer{
		Addr: kafka.TCP(brokers...),
	}
	defer w.Close()

	ctx := context.Background()
	if err := w.WriteMessages(ctx, messages...); err != nil {
		fmt.Printf("unable to write message: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	flag.StringVar(&file, "file", "", "the file contains messages")
	flag.Parse()

	if file == "" {
		panic("-file is a required flag")
	}
}
