package listLogs

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
var buildId int
var limit int

func ListLogInfo(cmd *cobra.Command, args []string) error {
	connection, ctx := azdoconfig.AzdoConfig()
	buildClient, _ := build.NewClient(ctx, connection)
	logRes, _ := buildClient.GetBuildLogs(ctx, build.GetBuildLogsArgs{Project: &project, BuildId: &buildId})
	if viper.GetInt("limit") > len(*logRes) {
		viper.Set("limit", len(*logRes))
	}
	for _, l := range (*logRes)[:viper.GetInt("limit")] {
		fmt.Fprintf(cmd.OutOrStdout(), "Log ID: "+strconv.Itoa(*l.Id)+" Line Count: "+strconv.Itoa(int(*l.LineCount))+" Created on: "+l.CreatedOn.String()+"\n")
	}
	return nil
}

func NewLogInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ListLogInfo",
		Short: "List Log info",
		Long:  "List log info for a build",
		RunE:  ListLogInfo,
		PreRun: func(cmd *cobra.Command, args []string) {
			err := viper.BindPFlag("limit", cmd.Flags().Lookup("limit"))
			if err != nil {
				log.Fatalln(errors.New("error binding limit flag"))
			}
		},
	}
	cmd.Flags().StringVarP(&project, "project", "p", "", "Project Name")
	cmd.Flags().IntVarP(&buildId, "buildId", "b", 0, "Build Id")
	cmd.Flags().IntVarP(&limit, "limit", "l", 10, "Limit")
	err := cmd.MarkFlagRequired("project")
	cmd.MarkFlagRequired("buildId")
	if err != nil {
		log.Fatal(errors.New("unable to mark project flag as required"))
	}

	return cmd
}
