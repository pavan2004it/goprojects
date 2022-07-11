package org

import (
	"context"
	"errors"
	"fmt"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"strconv"
)

type orgConfig struct {
	limit int
}

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
	if viper.GetInt("limit") > len(response.Value) {
		viper.Set("limit", len(response.Value))
	}
	for i, project := range response.Value[:viper.GetInt("limit")] {
		_, err2 := fmt.Fprintf(cmd.OutOrStdout(), strconv.Itoa(i)+" "+*project.Name+"\n")
		if err2 != nil {
			return err2
		}
	}
	return nil
}

func NewCmdOrg() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ListProjects",
		Short: "Lists All the projects in an Organization",
		Long:  "Calling all Org Api's in AZDO",
		RunE:  ListProjects,
		PreRun: func(cmd *cobra.Command, args []string) {
			err := viper.BindPFlag("limit", cmd.Flags().Lookup("limit"))
			if err != nil {
				log.Fatal(errors.New("error binding limit flag from org command"))
			}
		},
		Aliases: []string{"projectlist", "shprojects"},
	}
	orgConfig := &orgConfig{}
	cmd.Flags().IntVarP(&orgConfig.limit, "limit", "l", 5, "Result limit")

	return cmd
}
