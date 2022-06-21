//go:build integration

package utility_test

import (
	"testing"
	"utility"
)

func TestGetMultiCommandOutput(t *testing.T) {
	t.Parallel()
	text, err := utility.WaitMultipleCommand()
	if err != nil {
		t.Fatal(err)
	}
	commandOutput, err := utility.ParseCommandOutput(text)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(commandOutput.FirstCommand, commandOutput.SecondCommand)
}
