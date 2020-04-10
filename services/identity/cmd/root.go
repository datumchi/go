package cmd

import (
	"fmt"
	"github.com/datumchi/go/services/identity/configuration"
	"github.com/datumchi/go/utility/logger"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "identity",
	Short: "identity",
	Long: "DatumChi Identity Service",
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

	config := configuration.CreateConfiguration()

	logger.Infof("*******DatumChi Identity Service*******")
	logger.Infof("* SERVICE_HOST: %v", config.ServiceHost())
	logger.Infof("* SERVICE_PORT: %v", config.ServicePort())

}
