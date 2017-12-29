package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCmd())
}

func startCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "start",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	return cmd
}
