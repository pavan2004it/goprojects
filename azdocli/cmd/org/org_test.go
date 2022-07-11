package org_test

import (
	"azdocli/cmd"
	"azdocli/cmd/org"
	"bytes"
	"io/ioutil"
	"testing"
)

func TestRootCommand(t *testing.T) {
	t.Parallel()
	b := bytes.NewBufferString("")
	cmd.RootCmd.SetOut(b)
	cmd.RootCmd.SetArgs([]string{"-v"})
	err := cmd.RootCmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	want := "azdocli version: 0.0.1\n"
	got, err := ioutil.ReadAll(b)
	if want != string(got) {
		t.Fatalf("expected %q got %q", want, string(got))
	}

}

func TestListProjectsCommand(t *testing.T) {
	t.Parallel()
	projectCmd := org.NewCmdOrg()
	b := bytes.NewBufferString("")
	projectCmd.SetOut(b)
	projectCmd.SetArgs([]string{"-l", "2"})
	err := projectCmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	want := "0 GoProjects\n1 Docker\n"
	got, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != want {
		t.Fatalf("expected %q got %q", want, string(got))
	}

}
