package buildInfo_test

import (
	"azdocli/cmd"
	"azdocli/cmd/BuildInfo"
	"bytes"
	"io/ioutil"
	"testing"
)

func TestBuildInfoCmd(t *testing.T) {
	t.Parallel()
	rb := bytes.NewBufferString("")
	cmd.RootCmd.SetOut(rb)
	err := cmd.RootCmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	buildCmd := buildInfo.NewListBuildCmd()
	b := bytes.NewBufferString("")
	buildCmd.SetOut(b)
	buildCmd.SetArgs([]string{"-p", "Docker", "-l", "2"})
	err = buildCmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	want := "Build Id: 111 Build Name: Docker CI-CD Build Time: 2021-03-08 10:00:58.9689496 +0000 UTC\nBuild Id: 110 Build Name: Docker CI-CD Build Time: 2021-03-08 08:19:00.9052087 +0000 UTC\n"
	got, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != want {
		t.Fatalf("expected %q got %q", want, string(got))
	}
}
