package greet

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

type greetCon struct {
	greeting string
}

type option func(con *greetCon) error

func WithGreeting(greeting string) option {
	return func(con *greetCon) error {
		con.greeting = greeting
		return nil
	}
}

func NewGreetCon(opts ...option) *greetCon {
	greet := &greetCon{"Hello World"}
	for _, opt := range opts {
		err := opt(greet)
		if err != nil {
			log.Fatal(err)
		}
	}
	return greet
}

func Greet(cmd *cobra.Command, args []string) error {
	_, err := fmt.Fprintf(cmd.OutOrStdout(), viper.GetString("greeting")+"\n")
	if err != nil {
		return err
	}
	return nil
}

func NewGreetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "greet",
		Short: "Greet Commanmd",
		Long:  "Greet Command for testing",
		RunE:  Greet,
	}
	cmd.AddCommand(NewGreetCapsCommand())
	greeting := NewGreetCon(WithGreeting("Hello World"))
	cmd.PersistentFlags().StringVarP(&greeting.greeting, "greeting", "g", "Hello", "Greeting flag")
	bindErr := viper.BindPFlag("greeting", cmd.PersistentFlags().Lookup("greeting"))
	if bindErr != nil {
		log.Fatal(bindErr)
	}
	return cmd
}
