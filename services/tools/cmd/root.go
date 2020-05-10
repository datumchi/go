package cmd

import (
	"fmt"
	"github.com/datumchi/go/utility/logger"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "tools",
	Short: "tools",
	Long: "DatumChi Tools",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	cobra.OnInitialize(initConfig)
}

func initConfig() {

	logger.Infof("*******DatumChi Tools*******")

}
