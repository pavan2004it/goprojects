package buildInfo

import (
	"azdocli/pkg/azdoconfig"
	"errors"
	"fmt"
	"github.com/microsoft/azure-devops-go-api/azuredevops/build"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"strconv"
)

var project string
var limit int

func ListBuilds(cmd *cobra.Command, args []string) error {
	connection, ctx := azdoconfig.AzdoConfig()
	buildClient, _ := build.NewClient(ctx, connection)
	buildResponse, _ := buildClient.GetBuilds(ctx, build.GetBuildsArgs{Project: &project})
	if viper.GetInt("limit") > len(buildResponse.Value) {
		viper.Set("limit", len(buildResponse.Value))
	}
	for _, buildInfo := range buildResponse.Value[:viper.GetInt("limit")] {
		fmt.Fprintf(cmd.OutOrStdout(), "Build Id: "+strconv.Itoa(*buildInfo.Id)+" Build Name: "+*buildInfo.Definition.Name+" Build Time: "+buildInfo.FinishTime.String()+"\n")
	}
	return nil
}

func NewListBuildCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ListBuilds",
		Short: "List Builds",
		Long:  "List Builds for a project",
		RunE:  ListBuilds,
		PreRun: func(cmd *cobra.Command, args []string) {
			err := viper.BindPFlag("limit", cmd.Flags().Lookup("limit"))
			if err != nil {
				log.Fatalln(errors.New("error binding limit flag"))
			}
		},
	}
	cmd.Flags().StringVarP(&project, "project", "p", "", "Project Name")
	cmd.Flags().IntVarP(&limit, "limit", "l", 10, "Limit")
	err := cmd.MarkFlagRequired("project")
	if err != nil {
		log.Fatal(errors.New("unable to mark project flag as required"))
	}
	return cmd
}
