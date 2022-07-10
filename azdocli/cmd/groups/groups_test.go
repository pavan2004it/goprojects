package groups_test

import (
	"azdocli/cmd"
	"azdocli/cmd/groups"
	"bytes"
	"io/ioutil"
	"testing"
)

func TestListGroupCmd(t *testing.T) {
	t.Parallel()
	rb := bytes.NewBufferString("")
	cmd.RootCmd.SetOut(rb)
	err := cmd.RootCmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	b := bytes.NewBufferString("")
	groupCmd := groups.NewListGroupsCommand()
	groupCmd.SetOut(b)
	groupCmd.SetArgs([]string{"-l", "2"})
	err = groupCmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	want := "Docker Team\nProject Collection Build Administrators\n"
	got, err := ioutil.ReadAll(b)
	if want != string(got) {
		t.Fatalf("expected %q got %q", want, string(got))
	}
}

func TestListProjectSGCmd(t *testing.T) {
	t.Parallel()
	projectCmd := groups.NewProjectGroupsCommand()
	b := bytes.NewBufferString("")
	projectCmd.SetOut(b)
	projectCmd.SetArgs([]string{"projectsg", "-p", "Docker"})
	err := projectCmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	want := "Docker Team\nEndpoint Creators\n"
	got, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if want != string(got) {
		t.Fatalf("expected %q got %q", want, got)
	}
}
