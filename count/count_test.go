package count_test

import (
	"count"
	"testing"
)

//func TestLines(t *testing.T) {
//	t.Parallel()
//	inputBuf := bytes.NewBufferString("1\n2\n3")
//	c, err := count.NewCounter(
//		count.WithInput(inputBuf),
//	)
//	if err != nil {
//		t.Fatal(err)
//	}
//	want := 3
//	got := c.Lines()
//	if want != got {
//		t.Errorf("want %d, got %d", want, got)
//	}
//}

func TestWithInputFromArgs(t *testing.T) {
	t.Parallel()
	args := []string{"testdata/three_lines.txt"}
	c, err := count.NewCounter(
		count.WithInputFromMultiArgs(args),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 7
	got := c.MultiLines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestWithInputFromArgsEmpty(t *testing.T) {
	t.Parallel()
	c, err := count.NewCounter(
		count.WithInputFromMultiArgs([]string{}),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 0
	got := c.MultiLines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestWithInputFromMultiArgs(t *testing.T) {
	t.Parallel()
	args := []string{"testdata/three_lines.txt", "testdata/three_text.txt"}
	c, err := count.NewCounter(
		count.WithInputFromMultiArgs(args),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 12
	got := c.MultiLines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
