package stringFind_test

import (
	"github.com/google/go-cmp/cmp"
	"stringFind"
	"testing"
)

func TestStringFinder(t *testing.T) {
	t.Parallel()
	input := "sample.txt"
	l, err := stringFind.NewFinder(stringFind.WithInput(input), stringFind.WithSearch("Hello"))
	if err != nil {
		t.Fatal(err)
	}
	want := "Hello world\n" +
		"Hello, I am Gopher\n"
	got := l.Lines()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestSensitivity(t *testing.T) {
	t.Parallel()
	input := "sample.txt"
	l, err := stringFind.NewFinder(stringFind.WithInput(input), stringFind.WithSearch("heLlo"))
	if err != nil {
		t.Fatal(err)
	}

	want := "Hello world\n" +
		"Hello, I am Gopher\n"
	got := l.InSensitiveLines()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}
