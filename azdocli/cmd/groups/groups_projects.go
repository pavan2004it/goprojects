package groups

import (
	"context"
	"errors"
	"fmt"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	"github.com/microsoft/azure-devops-go-api/azuredevops/graph"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

type pConfig struct {
	project string
}

func ListGroupsInProjects(cmd *cobra.Command, args []string) error {
	configErr := viper.ReadInConfig()
	if configErr != nil {
		log.Fatal(configErr)
	}
	// Azdo Configuration
	organizationUrl := "https://dev.azure.com/" + viper.GetString("AZDO_ORG")
	personalAccessToken := viper.GetString("PAT_TOKEN")
	connection := azuredevops.NewPatConnection(organizationUrl, personalAccessToken)
	ctx := context.Background()

	// Client Configuration
	coreClient, err := core.NewClient(ctx, connection)
	groupClient, err := graph.NewClient(ctx, connection)
	if err != nil {
		return err
	}

	projects, _ := coreClient.GetProjects(ctx, core.GetProjectsArgs{})
	for _, p := range projects.Value {
		if *p.Name == viper.GetString("project") {
			desRes, _ := groupClient.GetDescriptor(ctx, graph.GetDescriptorArgs{StorageKey: p.Id})
			groups, err := groupClient.ListGroups(ctx, graph.ListGroupsArgs{ScopeDescriptor: desRes.Value})
			if err != nil {
				log.Fatal(err)
			}
			if viper.GetInt("limit") > len(*groups.GraphGroups) {
				viper.Set("limit", len(*groups.GraphGroups))
			}
			for _, group := range (*groups.GraphGroups)[:viper.GetInt("limit")] {
				_, err2 := fmt.Fprintf(cmd.OutOrStdout(), *group.DisplayName+"\n")
				if err2 != nil {
					return err2
				}
			}
		}
	}
	return nil
}

func NewProjectGroupsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "projectsg",
		Short: "List Security Groups in a Project",
		Long:  "Lists Security Groups for a project in an organization",
		RunE:  ListGroupsInProjects,
		PreRun: func(cmd *cobra.Command, args []string) {
			err := viper.BindPFlag("project", cmd.Flags().Lookup("project"))
			if err != nil {
				log.Fatal(errors.New("failed to bind flag project in the command projectsg"))
			}
		},
		Aliases: []string{"shprojsg", "psg"},
	}
	projectConfig := pConfig{}
	cmd.Flags().StringVarP(&projectConfig.project, "project", "p", "", "project name")
	err := cmd.MarkFlagRequired("project")
	if err != nil {
		log.Fatal(errors.New("failed to mark flag project as required"))
	}
	return cmd
}
