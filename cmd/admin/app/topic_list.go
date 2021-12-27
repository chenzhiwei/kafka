package app

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	topicListCmd = &cobra.Command{
		Use:   "list",
		Short: "list all topics",
		RunE: func(_ *cobra.Command, _ []string) error {
			if err := listTopics(); err != nil {
				return err
			}
			return nil
		},
	}
)

func listTopics() error {
	conn, err := brokerConfig.Connection()
	if err != nil {
		fmt.Printf("unable to make connection to broker: %s, error: %v\n", brokerConfig.Broker, err)
		os.Exit(1)
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		fmt.Printf("unable to read partitions: %v\n", err)
		os.Exit(1)
	}

	topics := map[string]struct{}{}
	for _, p := range partitions {
		topics[p.Topic] = struct{}{}
	}

	for topic := range topics {
		fmt.Printf("%s\n", topic)
	}

	return nil
}
