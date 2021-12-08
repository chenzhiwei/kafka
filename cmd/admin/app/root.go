package app

import (
	"github.com/spf13/cobra"
)

var (
	broker string

	rootCmd = &cobra.Command{
		Use:          "kafka-admin",
		Short:        "kafka-admin is a tool to manage Kafka",
		SilenceUsage: true,
	}
)

func init() {
	rootCmd.AddCommand(topicCmd)
	rootCmd.PersistentFlags().StringVar(&broker, "broker", "localhost:9092", "the broker address")
}

func Execute() error {
	return rootCmd.Execute()
}
