package usermembership

import (
	"context"
	"errors"
	"fmt"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/graph"
	"github.com/microsoft/azure-devops-go-api/azuredevops/memberentitlementmanagement"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"regexp"
)

type userConfig struct {
	username string
}

func ListUserSg(cmd *cobra.Command, args []string) error {
	sg := map[string][]string{}
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	organizationUrl := "https://dev.azure.com/" + viper.GetString("AZDO_ORG")
	personalAccessToken := viper.GetString("PAT_TOKEN")
	connection := azuredevops.NewPatConnection(organizationUrl, personalAccessToken)
	ctx := context.Background()
	userClient, _ := memberentitlementmanagement.NewClient(ctx, connection)
	groupClient, err := graph.NewClient(ctx, connection)
	members, _ := userClient.GetUserEntitlements(ctx, memberentitlementmanagement.GetUserEntitlementsArgs{})
	for _, member := range *members.Members {
		if *member.User.PrincipalName == viper.GetString("user") {
			fmt.Fprintf(cmd.OutOrStdout(), "\n"+*member.User.DisplayName+" has access to below groups in respective projects:"+"\n")
			groups, err := groupClient.ListMemberships(ctx, graph.ListMembershipsArgs{SubjectDescriptor: member.User.Descriptor})
			if err != nil {
				log.Fatal(err)
			}
			for _, group := range *groups {
				groupDetails, _ := groupClient.GetGroup(ctx, graph.GetGroupArgs{GroupDescriptor: group.ContainerDescriptor})
				RegexLogic(*groupDetails.PrincipalName, sg, *groupDetails.DisplayName)
			}
			for key, value := range sg {
				fmt.Fprintf(cmd.OutOrStdout(), "\nProject: "+key+"\n\n")
				fmt.Fprintf(cmd.OutOrStdout(), "Groups: "+"\n")
				for _, v := range value {
					fmt.Fprintf(cmd.OutOrStdout(), v+"\n")
				}
			}
			sg = map[string][]string{}
			return nil
		}
	}
	return errors.New("user not found")
}

func RegexLogic(key string, data map[string][]string, value string) map[string][]string {
	re := regexp.MustCompile(`\[(.*)\]`)
	pattern := re.FindStringSubmatch(key)
	_, ok := data[pattern[1]]
	if ok {
		data[pattern[1]] = append(data[pattern[1]], value)
	} else {
		data[pattern[1]] = []string{value}
	}
	return data
}

func NewListUserSgCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ListUserSgs",
		Short: "List Security Groups for a user",
		Long:  "Lists Security Groups for an Organization",
		RunE:  ListUserSg,
		PreRun: func(cmd *cobra.Command, args []string) {
			err := viper.BindPFlag("user", cmd.Flags().Lookup("user"))
			if err != nil {
				log.Fatal(errors.New("unable to bind flag user in ListUserSgs"), err)
			}
		},
		Aliases: []string{"showaccess"},
	}
	userCfg := userConfig{}
	cmd.Flags().StringVarP(&userCfg.username, "user", "u", "", "username")
	cmd.MarkFlagRequired("user")
	return cmd
}
