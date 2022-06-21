package utility_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"os"
	"testing"
	"utility"
)

func TestParseCommandOutput(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile("testdata/sample.txt")
	if err != nil {
		t.Fatal(err)
	}
	want := utility.CommandOutput{FirstCommand: "27952", SecondCommand: "utility_test.go"}

	got, err := utility.ParseCommandOutput(string(data))
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(utility.CommandOutput{})) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestWaitMultipleCommand(t *testing.T) {
	t.Parallel()
	data, err := utility.WaitMultipleCommand()
	if err != nil {
		t.Fatal(err)
	}
	want := utility.CommandOutput{FirstCommand: "27953", SecondCommand: "utility_test.go"}

	got, err := utility.ParseCommandOutput(string(data))
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(utility.CommandOutput{})) {
		t.Error(cmp.Diff(want, got))
	}

}

func TestCommandTimeout(t *testing.T) {
	t.Parallel()
	err := utility.CommandTimeout()
	if err == nil {
		t.Errorf("wanted error %v, but got nil", err)
	}
}
