package cmd

import (
	"azdocli/cmd/org"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "azdocli",
	Short: "Cli to interact with Azure DevOps API's",
	Long:  "Cli to run azdo commands",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "provide the path for config file else it will be searched in the home directory")
	rootCmd.AddCommand(org.CmdOrg)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.SetVersionTemplate("azdocli version: {{.Version}}\n")
	rootCmd.Version = "0.0.1"

}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name "app".
		viper.AddConfigPath(home)
		viper.SetConfigName("app")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}