package cmd

import (
	"azdocli/cmd/org"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

type orgConfig struct {
	orgName string
}

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
	orgConfig := &orgConfig{orgName: ""}
	rootCmd.AddCommand(org.CmdOrg)
	rootCmd.PersistentFlags().StringVarP(&orgConfig.orgName, "org", "o", "", "org name")
	viper.BindPFlag("org", rootCmd.PersistentFlags().Lookup("org"))
	viper.AddConfigPath("/Users/pavankumar/goprojects/azdocli/config/")
	viper.SetConfigName("app")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
