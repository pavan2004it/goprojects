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
	orgName  string
	username string
}

func ListUserSg(cmd *cobra.Command, args []string) error {
	sg := map[string][]string{}
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	organizationUrl := "https://dev.azure.com/" + viper.GetString("orz")
	personalAccessToken := viper.GetString("PAT_TOKEN")
	connection := azuredevops.NewPatConnection(organizationUrl, personalAccessToken)
	ctx := context.Background()
	userClient, _ := memberentitlementmanagement.NewClient(ctx, connection)
	groupClient, err := graph.NewClient(ctx, connection)
	members, _ := userClient.GetUserEntitlements(ctx, memberentitlementmanagement.GetUserEntitlementsArgs{})
	for _, member := range *members.Members {
		if *member.User.PrincipalName == viper.GetString("user") {
			fmt.Println("\n" + *member.User.DisplayName + " has access to below groups in respective projects:")
			groups, err := groupClient.ListMemberships(ctx, graph.ListMembershipsArgs{SubjectDescriptor: member.User.Descriptor})
			if err != nil {
				log.Fatal(err)
			}
			for _, group := range *groups {
				groupDetails, _ := groupClient.GetGroup(ctx, graph.GetGroupArgs{GroupDescriptor: group.ContainerDescriptor})
				RegexLogic(*groupDetails.PrincipalName, sg, *groupDetails.DisplayName)
			}
			OutData(&sg)
		}
	}
	for _, member := range *members.Members {
		if *member.User.PrincipalName != viper.GetString("user") {
			return errors.New("user not found")
		}
	}
	return nil
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

func OutData(data *map[string][]string) {
	cmd := NewListUserSgCommand()
	for key, value := range *data {
		fmt.Fprintf(cmd.OutOrStdout(), "\nProject: "+key+"\n\n")
		fmt.Fprintf(cmd.OutOrStdout(), "Groups: "+"\n")
		for _, v := range value {
			fmt.Println(v)
		}
	}
	*data = map[string][]string{}
}

func NewListUserSgCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "usersg",
		Short: "List Security Groups",
		Long:  "Lists Security Groups for an Organization",
		RunE:  ListUserSg,
		PreRun: func(cmd *cobra.Command, args []string) {
			err := viper.BindPFlag("orz", cmd.Flags().Lookup("org"))
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	userCfg := userConfig{}
	cmd.Flags().StringVarP(&userCfg.orgName, "org", "o", "", "org name")
	cmd.Flags().StringVarP(&userCfg.username, "user", "u", "", "username")
	cmd.MarkFlagRequired("user")
	cmd.MarkFlagRequired("orz")
	err := viper.BindPFlag("user", cmd.Flags().Lookup("user"))
	if err != nil {
		log.Fatal(err)
	}
	return cmd
}
