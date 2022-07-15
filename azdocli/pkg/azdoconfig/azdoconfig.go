package azdoconfig

import (
	"context"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/spf13/viper"
	"log"
)

func AzdoConfig() (*azuredevops.Connection, context.Context) {
	configErr := viper.ReadInConfig()
	if configErr != nil {
		log.Fatal(configErr)
	}
	organizationUrl := "https://dev.azure.com/" + viper.GetString("AZDO_ORG")
	personalAccessToken := viper.GetString("PAT_TOKEN")
	connection := azuredevops.NewPatConnection(organizationUrl, personalAccessToken)
	ctx := context.Background()
	return connection, ctx
}
