package getlog

import (
	"azdocli/pkg/azdoconfig"
	"bufio"
	"fmt"
	"github.com/microsoft/azure-devops-go-api/azuredevops/build"
	"github.com/spf13/cobra"
	"strings"
)

var project string
var buildId int
var logId int
var pattern string

func GetBuildLog(cmd *cobra.Command, args []string) error {
	connection, ctx := azdoconfig.AzdoConfig()
	buildClient, _ := build.NewClient(ctx, connection)
	res, _ := buildClient.GetBuildLog(ctx, build.GetBuildLogArgs{Project: &project, BuildId: &buildId, LogId: &logId})
	scanner := bufio.NewScanner(res)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, cmd.Flags().Lookup("match").Value.String()) {
			fmt.Fprintf(cmd.OutOrStdout(), line+"\n")
		}
	}
	return nil
}

func NewBuildLogCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "GetBuildLog",
		Short: "Get Build Log",
		Long:  "Get Build Log",
		RunE:  GetBuildLog,
	}
	cmd.Flags().StringVarP(&project, "project", "p", "", "Project Name")
	cmd.Flags().IntVarP(&buildId, "buildid", "b", 0, "Build Id")
	cmd.Flags().IntVarP(&logId, "logid", "l", 0, "Log Id")
	cmd.Flags().StringVarP(&pattern, "match", "m", "", "Match pattern")
	return cmd
}
