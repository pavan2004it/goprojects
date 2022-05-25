package myTypes_test

import (
	"myTypes"
	"strings"
	"testing"
)

func TestTwice(t *testing.T) {
	t.Parallel()
	input := myTypes.MyInt(9)
	want := myTypes.MyInt(18)
	got := input.Twice()

	if want != got {
		t.Errorf("twice %d: want %d, got %d",
			input, want, got)
	}
}
func TestMyStringLen(t *testing.T) {
	t.Parallel()
	input := myTypes.MyString("hello")
	want := 5
	got := input.MyStringLen()

	if want != got {
		t.Errorf("len %s: want %d, got %d",
			input, want, got)
	}
}

// test case to test the string builder functionality

func TestStringsBuilder(t *testing.T) {
	t.Parallel()
	var sb strings.Builder
	sb.WriteString("hello, ")
	sb.WriteString("Gophers!")
	want := "hello, Gophers!"
	got := sb.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
	wantLen := 15
	gotLen := sb.Len()
	if wantLen != gotLen {
		t.Errorf("%q: want len %d, got %d",
			sb.String(), wantLen, gotLen)
	}
}

// test case to test the custom type functionality

func TestMyBuilderHello(t *testing.T) {
	t.Parallel()
	var mb myTypes.MyBuilder
	mb.WriteString("Hello, ")
	mb.WriteString("Gophers!")
	want := "Hello, Gophers!"
	got := mb.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
	wantLen := 15
	gotLen := mb.Len()
	if wantLen != gotLen {
		t.Errorf("%q: want len %d, got %d",
			mb.String(), wantLen, gotLen)
	}
}

func TestStringUppercaser(t *testing.T) {
	t.Parallel()
	var su myTypes.StringUppercaser
	su.Contents.WriteString("Hello, Gophers!")
	want := "HELLO, GOPHERS!"
	got := su.ToUpper()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestDouble(t *testing.T) {
	t.Parallel()
	x := myTypes.MyInt(12)
	want := myTypes.MyInt(24)
	p := &x
	p.Double()
	if want != x {
		t.Errorf("want %d, got %d", want, x)
	}
}
