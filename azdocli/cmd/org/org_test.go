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

func TestProjectCommand(t *testing.T) {
	t.Parallel()
	projectCmd := org.NewCmdOrg()
	b := bytes.NewBufferString("")
	projectCmd.SetOut(b)
	projectCmd.SetArgs([]string{""})
	err := projectCmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	want := "0 Docker\n1 InfrastructureAsCode\n2 MyFirstProject\n3 gameoflife\n"
	got, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != want {
		t.Fatalf("expected %q got %q", want, string(got))
	}

}

func TestUserCommand(t *testing.T) {
	t.Parallel()
	userCmd := org.NewUserCmd()
	b := bytes.NewBufferString("")
	userCmd.SetOut(b)
	userCmd.SetArgs([]string{"users"})
	err := userCmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	want := "glaringfireball@gmail.com\npavan_2004it@hotmail.com\npavan2004it@gmail.com\n"
	got, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != want {
		t.Fatalf("expected %q got %q", want, string(got))
	}
}
