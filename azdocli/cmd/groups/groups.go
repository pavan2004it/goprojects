package groups

import (
	"context"
	"fmt"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/graph"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

type groupConfig struct {
	limit   int
	orgName string
}

func ListGroups(cmd *cobra.Command, args []string) error {
	configErr := viper.ReadInConfig()
	if configErr != nil {
		log.Fatal(configErr)
	}
	organizationUrl := "https://dev.azure.com/" + viper.GetString("org") // todo: replace value with your organization url
	personalAccessToken := viper.GetString("PAT_TOKEN")
	connection := azuredevops.NewPatConnection(organizationUrl, personalAccessToken)
	ctx := context.Background()
	groupClient, err := graph.NewClient(ctx, connection)
	if err != nil {
		return err
	}
	response, err := groupClient.ListGroups(ctx, graph.ListGroupsArgs{})
	if err != nil {
		log.Fatal(err)
	}

	for _, group := range (*response.GraphGroups)[:viper.GetInt("limit")] {
		fmt.Fprintf(cmd.OutOrStdout(), *group.DisplayName+"\n")
	}

	return nil
}

func NewListGroupsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "groupsg",
		Short: "List Security Groups",
		Long:  "Lists Security Groups for an Organization",
		RunE:  ListGroups,
		PreRun: func(cmd *cobra.Command, args []string) {
			orgErr := viper.BindPFlag("org", cmd.PersistentFlags().Lookup("org"))
			if orgErr != nil {
				log.Fatal(orgErr)
			}
		},
	}
	gConfig := groupConfig{}
	cmd.AddCommand(NewProjectGroupsCommand())
	cmd.PersistentFlags().StringVarP(&gConfig.orgName, "org", "o", "", "org name")
	cmd.PersistentFlags().IntVarP(&gConfig.limit, "limit", "l", 0, "Result limit")
	err := viper.BindPFlag("limit", cmd.PersistentFlags().Lookup("limit"))
	if err != nil {
		return nil
	}
	cmd.MarkPersistentFlagRequired("org")
	return cmd
}
