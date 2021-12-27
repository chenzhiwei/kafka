package app

import (
	"fmt"
	"os"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
)

var (
	partition int
	replicas  int

	topicCreateCmd = &cobra.Command{
		Use:   "create topic1 topic2 topic3",
		Short: "Create one or multiple topics",
		RunE: func(_ *cobra.Command, args []string) error {
			if err := createTopics(args); err != nil {
				return err
			}
			return nil
		},
	}
)

func init() {
	topicCreateCmd.Flags().IntVar(&partition, "partition", 1, "the partition number of the topic")
	topicCreateCmd.Flags().IntVar(&replicas, "replicas", 1, "the number of replicas of each partition")
}

func createTopics(topics []string) error {
	conn, err := brokerConfig.Connection()
	if err != nil {
		fmt.Printf("unable to make connection to broker: %s, error: %v\n", brokerConfig.Broker, err)
		os.Exit(1)
	}
	defer conn.Close()

	for _, topic := range topics {
		topicConfig := kafka.TopicConfig{
			Topic:             topic,
			NumPartitions:     partition,
			ReplicationFactor: replicas,
		}

		if err := conn.CreateTopics(topicConfig); err != nil {
			fmt.Printf("unable to create topic: %s, error: %v\n", topic, err)
		} else {
			fmt.Printf("topic %s created successfully\n", topic)
		}
	}

	return nil
}
