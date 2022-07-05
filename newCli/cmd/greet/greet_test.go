package greet_test

import (
	"bytes"
	"io/ioutil"
	"newCli/cmd/greet"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Parallel()
	greetCmd := greet.NewGreetCommand()
	b := bytes.NewBufferString("")
	greetCmd.SetOut(b)
	greetCmd.SetArgs([]string{"-g", "Hello"})
	err := greetCmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	want := "Hello\n"
	got, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != want {
		t.Fatalf("expected %q got %q", want, string(got))
	}

}

func TestGreetCaps(t *testing.T) {
	t.Parallel()
	greetCapsCmd := greet.NewGreetCapsCommand()
	b := bytes.NewBufferString("")
	greetCapsCmd.SetOut(b)
	greetCapsCmd.SetArgs([]string{"caps"})
	err := greetCapsCmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	want := "HELLO\n"
	got, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != want {
		t.Fatalf("expected %q got %q", want, string(got))
	}
}
