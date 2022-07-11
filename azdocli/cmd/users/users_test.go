package users_test

import (
	"azdocli/cmd"
	"azdocli/cmd/users"
	"bytes"
	"io/ioutil"
	"testing"
)

func TestListUserSgsCommand(t *testing.T) {
	t.Parallel()
	rb := bytes.NewBufferString("")
	cmd.RootCmd.SetOut(rb)
	err := cmd.RootCmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	userCmd := users.NewUserCmd()
	b := bytes.NewBufferString("")
	userCmd.SetOut(b)
	userCmd.SetArgs([]string{""})
	err = userCmd.Execute()
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
