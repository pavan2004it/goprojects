package listLogs_test

import (
	"azdocli/cmd"
	"azdocli/cmd/listLogs"
	"bytes"
	"io/ioutil"
	"testing"
)

func TestLogInfoCmd(t *testing.T) {
	t.Parallel()
	rb := bytes.NewBufferString("")
	cmd.RootCmd.SetOut(rb)
	err := cmd.RootCmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	logCmd := listLogs.NewLogInfoCmd()
	b := bytes.NewBufferString("")
	logCmd.SetOut(b)
	logCmd.SetArgs([]string{"-p", "Docker", "-b", "1", "-l", "2"})
	err = logCmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	want := "Log ID: 3 Line Count: 24 Created on: 2021-03-02 12:48:28.94 +0000 UTC\nLog ID: 4 Line Count: 27 Created on: 2021-03-02 12:48:45.83 +0000 UTC\n"
	got, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != want {
		t.Fatalf("expected %q got %q", want, string(got))
	}
}
