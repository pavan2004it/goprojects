package groups_test

import (
	"azdocli/cmd/groups"
	"bytes"
	"io/ioutil"
	"testing"
)

func TestListGroupCmd(t *testing.T) {
	t.Parallel()
	b := bytes.NewBufferString("")
	groupCmd := groups.NewListGroupsCommand()
	groupCmd.SetOut(b)
	groupCmd.SetArgs([]string{"-o", "pavantikkani"})
	err := groupCmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	want := "Project Administrator"
	got, err := ioutil.ReadAll(b)
	if want != string(got) {
		t.Fatalf("expected %q got %q", want, string(got))
	}
}
