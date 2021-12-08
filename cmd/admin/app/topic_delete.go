package app

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/segmentio/kafka-go"
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
	conn, err := kafka.Dial("tcp", broker)
	if err != nil {
		fmt.Printf("unable to dial broker: %s, error: %v\n", broker, err)
		os.Exit(1)
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		fmt.Printf("unable to request controller: %v\n", err)
		os.Exit(1)
	}
	ctrConn, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		fmt.Printf("unable to dial controller: %v\n", err)
		os.Exit(1)
	}
	defer ctrConn.Close()

	for _, topic := range topics {
		if err := ctrConn.DeleteTopics(topic); err != nil {
			fmt.Printf("unable to delete topic: %s, error: %v\n", topic, err)
		} else {
			fmt.Printf("topic %s deleted successfully\n", topic)
		}
	}

	return nil
}
