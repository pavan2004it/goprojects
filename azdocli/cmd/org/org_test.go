package org_test

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestProjectList(t *testing.T) {
	t.Parallel()
	out, err := exec.Command("azdocli", "projects", "-o", "pavantikkani").CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}
	want := "0 Docker\n1 InfrastructureAsCode\n2 MyFirstProject\n3 gameoflife\n"
	got := bytes.NewBuffer(out)
	if want != got.String() {
		t.Errorf("Wanted %q got %q", want, got)
	}
}

func TestUserList(t *testing.T) {
	t.Parallel()
	out, err := exec.Command("azdocli", "projects", "-o", "pavantikkani", "users").CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}
	want := "glaringfireball@gmail.com\npavan_2004it@hotmail.com\npavan2004it@gmail.com\n"
	got := bytes.NewBuffer(out)
	if want != got.String() {
		t.Errorf("Wanted %q got %q", want, got)
	}
}
