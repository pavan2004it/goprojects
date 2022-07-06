package org

import (
	"context"
	"fmt"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/memberentitlementmanagement"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

func ListUsers(cmd *cobra.Command, args []string) error {

	configErr := viper.ReadInConfig()
	if configErr != nil {
		log.Fatal(configErr)
	}
	organizationUrl := "https://dev.azure.com/" + viper.GetString("org") // todo: replace value with your organization url
	personalAccessToken := viper.GetString("PAT_TOKEN")
	connection := azuredevops.NewPatConnection(organizationUrl, personalAccessToken)
	ctx := context.Background()
	memberClient, _ := memberentitlementmanagement.NewClient(ctx, connection)
	response, _ := memberClient.GetUserEntitlements(ctx, memberentitlementmanagement.GetUserEntitlementsArgs{})
	for _, member := range *response.Members {
		_, err := fmt.Fprintf(cmd.OutOrStdout(), *member.User.PrincipalName+"\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func NewUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "users",
		Short: "Lists All Users in the organization",
		Long:  "Calling User entitlement API",
		RunE:  ListUsers,
	}
	return cmd
}
