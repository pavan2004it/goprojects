package greet

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

func CapsGreet(cmd *cobra.Command, args []string) error {
	_, err := fmt.Fprintf(cmd.OutOrStdout(), strings.ToUpper(viper.GetString("greeting"))+"\n")
	if err != nil {
		return err
	}
	return nil
}

func NewGreetCapsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "caps",
		Short: "Capitalizes Greeting",
		Long:  "Caps command capitalizes greeting",
		RunE:  CapsGreet,
	}
	return cmd
}
