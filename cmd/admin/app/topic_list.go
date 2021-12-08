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

	partitions, err := ctrConn.ReadPartitions()
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
