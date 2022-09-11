package cmd

import (
	"fmt"
	"os"

	"github.com/bcpitutor/tiki/appconfig"
	"github.com/bcpitutor/tiki/utils"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var TLogger *zap.Logger

var rootCmd = &cobra.Command{
	Use:   "tiki",
	Short: "",
	Long:  ``,
}

func init() {
	rootCmd.PersistentFlags().StringP("profile", "p", "default", "Profile to use")
	rootCmd.PersistentFlags().BoolP("json", "j", false, "JSON Output")

	tikidir := utils.InitalizeTikiDirectory()
	err := appconfig.InitConfig(tikidir)
	if err != nil {
		fmt.Printf("Cannot read config file: %s\n", err)
		os.Exit(-1)
	}

	// profile := rootCmd.Flag("profile").Value.String()
	// fmt.Printf("Setting %s into selected profile\n", profile)
	//appconfig.AppConfig.SelectedProfile = profile
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
