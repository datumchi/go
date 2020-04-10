package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the identity service",
	Long:  "Starts the identity service",
	Run: func(cmd *cobra.Command, args []string) {
		// service.Start()
	},
}
