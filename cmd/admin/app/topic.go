package app

import "github.com/spf13/cobra"

var topicCmd = &cobra.Command{
	Use:   "topic create|delete|list|describe [ARGS]",
	Short: "create, delete, list, describe topics",
}

func init() {
	topicCmd.AddCommand(topicCreateCmd)
	topicCmd.AddCommand(topicDeleteCmd)
	topicCmd.AddCommand(topicListCmd)
}
