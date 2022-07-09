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
	"strconv"
)

func ListProjects(cmd *cobra.Command, args []string) error {
	configErr := viper.ReadInConfig()
	if configErr != nil {
		log.Fatal(configErr)
	}
	organizationUrl := "https://dev.azure.com/" + viper.GetString("AZDO_ORG") // todo: replace value with your organization url
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
		_, err2 := fmt.Fprintf(cmd.OutOrStdout(), strconv.Itoa(i)+" "+*project.Name+"\n")
		if err2 != nil {
			return err2
		}
	}
	return nil
}

func NewCmdOrg() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "projects",
		Short: "Lists All the projects in an Organization",
		Long:  "Calling all Org Api's in AZDO",
		RunE:  ListProjects,
	}
	cmd.AddCommand(NewUserCmd())
	return cmd
}
