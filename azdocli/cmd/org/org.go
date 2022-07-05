package org

import "C"
import (
	"context"
	"fmt"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

type orgConfig struct {
	orgName string
}

var CmdOrg = &cobra.Command{
	Use:   "projects",
	Short: "Lists All the projects in an Organization",
	Long:  "Calling all Org Api's in AZDO",
	Run: func(cmd *cobra.Command, args []string) {

		//viper.SetConfigFile("app.env")
		configErr := viper.ReadInConfig()
		if configErr != nil {
			log.Fatal(configErr)
		}
		organizationUrl := "https://dev.azure.com/" + viper.GetString("org") // todo: replace value with your organization url
		personalAccessToken := viper.GetString("PAT_TOKEN")
		connection := azuredevops.NewPatConnection(organizationUrl, personalAccessToken)
		ctx := context.Background()

		coreClient, err := core.NewClient(ctx, connection)
		if err != nil {
			log.Fatal(err)
		}
		response, err := coreClient.GetProjects(ctx, core.GetProjectsArgs{})
		if err != nil {
			log.Fatal(err)
		}
		for i, project := range response.Value {
			fmt.Println(i, *project.Name)
		}

	},
}

func init() {
	orgConfig := &orgConfig{orgName: ""}
	CmdOrg.PersistentFlags().StringVarP(&orgConfig.orgName, "org", "o", "", "org name")
	bindErr := viper.BindPFlag("org", CmdOrg.PersistentFlags().Lookup("org"))
	if bindErr != nil {
		log.Fatal(bindErr)
	}
}
