package cmd

import (
	"github.com/datumchi/go/services/hborderer/service"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the hborderer service",
	Long:  "Starts the hborderer service",
	Run: func(cmd *cobra.Command, args []string) {
		service.Start()
	},
}
