package boolf

import (
	"fmt"
	"github.com/spf13/cobra"
)

func Bool(cmd *cobra.Command, args []string) error {
	flagVal, _ := cmd.Flags().GetBool("boolf")
	if flagVal {
		withbool()
	} else {
		withoutbool()
	}
	return nil
}

func NewBoolCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "boolf",
		Short: "Bool Command",
		Long:  "Bool Command for testing",
		RunE:  Bool,
	}

	cmd.Flags().BoolP("boolf", "b", false, "Bool flag")
	return cmd

}

func withbool() {
	fmt.Println("Bool Declared")
}

func withoutbool() {
	fmt.Println("Bool Not Declared")
}
