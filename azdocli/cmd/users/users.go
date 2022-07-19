package users

import (
	"azdocli/pkg/azdoconfig"
	"fmt"
	"github.com/microsoft/azure-devops-go-api/azuredevops/memberentitlementmanagement"
	"github.com/spf13/cobra"
)

func ListUsers(cmd *cobra.Command, args []string) error {

	connection, ctx := azdoconfig.AzdoConfig()
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
		Use:     "ListUsers",
		Short:   "Lists All Users in the organization",
		Long:    "Calling User entitlement API",
		RunE:    ListUsers,
		Aliases: []string{"userlist", "shusers"},
	}
	return cmd
}
