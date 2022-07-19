package getlog_test

import (
	"azdocli/cmd"
	"azdocli/cmd/getlog"
	"bytes"
	"io/ioutil"
	"testing"
)

func TestNewBuildLogCmd(t *testing.T) {
	t.Parallel()
	rb := bytes.NewBufferString("")
	cmd.RootCmd.SetOut(rb)
	err := cmd.RootCmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	buildLogCmd := getlog.NewBuildLogCmd()
	b := bytes.NewBufferString("")
	buildLogCmd.SetOut(b)
	buildLogCmd.SetArgs([]string{"-p", "Docker", "-b", "111", "-l", "2", "-m", "Greatest"})
	err = buildLogCmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	want := "Greatest Parser Depth: 12 (Max: 100)\nGreatest File Size: 1,609 (Max: 1,048,576)\n"
	got, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != want {
		t.Fatalf("expected %q got %q", want, string(got))
	}
}
