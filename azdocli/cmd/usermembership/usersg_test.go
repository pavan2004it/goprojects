package usermembership_test

import (
	"azdocli/cmd"
	"azdocli/cmd/usermembership"
	"bytes"
	"io/ioutil"
	"testing"
)

func TestUserListSGCmd(t *testing.T) {
	t.Parallel()
	rb := bytes.NewBufferString("")
	cmd.RootCmd.SetOut(rb)
	err := cmd.RootCmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	userCmd := usermembership.NewListUserSgCommand()
	b := bytes.NewBufferString("")
	userCmd.SetOut(b)
	userCmd.SetArgs([]string{"-u", "pavan2004it@gmail.com"})
	err = userCmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	want := "\npavan2004it@gmail.com has access to below groups in respective projects:\n\nProject: Docker\n\nGroups: \nContributors\n"
	got, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if want != string(got) {
		t.Fatalf("expected %q got %q", want, got)
	}

}
