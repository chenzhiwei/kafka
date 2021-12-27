package app

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	topicDeleteCmd = &cobra.Command{
		Use:   "delete topic1 topic2 topic3",
		Short: "delete one or multiple topics",
		RunE: func(_ *cobra.Command, args []string) error {
			if err := deleteTopics(args); err != nil {
				return err
			}
			return nil
		},
	}
)

func deleteTopics(topics []string) error {
	conn, err := brokerConfig.Connection()
	if err != nil {
		fmt.Printf("unable to make connection to broker: %s, error: %v\n", brokerConfig.Broker, err)
		os.Exit(1)
	}
	defer conn.Close()

	for _, topic := range topics {
		if err := conn.DeleteTopics(topic); err != nil {
			fmt.Printf("unable to delete topic: %s, error: %v\n", topic, err)
		} else {
			fmt.Printf("topic %s deleted successfully\n", topic)
		}
	}

	return nil
}
