package app

import (
	"github.com/spf13/cobra"

	"github.com/chenzhiwei/kafka/utils/config"
)

var (
	brokerConfig config.Config

	rootCmd = &cobra.Command{
		Use:          "kafka-admin",
		Short:        "kafka-admin is a tool to manage Kafka",
		SilenceUsage: true,
	}
)

func init() {
	rootCmd.AddCommand(topicCmd)
	rootCmd.PersistentFlags().StringVar(&brokerConfig.Broker, "broker", "localhost:9092", "the broker address")
	rootCmd.PersistentFlags().BoolVar(&brokerConfig.SkipTLSVerify, "skip-tls-verify", false, "skip TLS verification")
	rootCmd.PersistentFlags().StringVar(&brokerConfig.CaFile, "ca-file", "", "CA file for tls connection")
	rootCmd.PersistentFlags().StringVar(&brokerConfig.ClientKeyFile, "client-key-file", "", "Client key file for mTLS connection")
	rootCmd.PersistentFlags().StringVar(&brokerConfig.ClientCertFile, "client-cert-file", "", "Client cert file for mTLS connection")
	rootCmd.PersistentFlags().StringVar(&brokerConfig.Mechanism, "mechanism", "", "SASL authentication mechanism")
	rootCmd.PersistentFlags().StringVar(&brokerConfig.Username, "username", "", "the login username")
	rootCmd.PersistentFlags().StringVar(&brokerConfig.Password, "password", "", "the login password")
}

func Execute() error {
	return rootCmd.Execute()
}
