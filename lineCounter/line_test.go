package lineCounter_test

import (
	"bytes"
	"io"
	"lineCounter"
	"testing"
)

func TestLines(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("\"1\n2\n3\"")
	c, err := lineCounter.NewCounter(lineCounter.WithInput(inputBuf))
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Lines()

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}

}

func TestWithInputFromArgs(t *testing.T) {
	t.Parallel()
	args := []string{"testdata/three_lines.txt"}
	c, err := lineCounter.NewCounter(lineCounter.FromArgs(args))
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Lines()

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}

}

func TestWithInputFromArgsEmpty(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("1\n2\n3")
	c, err := lineCounter.NewCounter(
		lineCounter.WithInput(inputBuf),
		lineCounter.FromArgs([]string{}),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestWords(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("1\n2 words\n3 this time")
	c, err := lineCounter.NewCounter(lineCounter.WithInput(inputBuf))
	if err != nil {
		t.Fatal(err)
	}
	want := 6
	got := c.Words()

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}

}

func TestWordCount(t *testing.T) {
	t.Parallel()
	args := []string{"-w", "testdata/three_lines.txt"}
	c, err := lineCounter.NewCounter(
		lineCounter.FromArgs(args),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 6
	got := c.Words()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestFromArgsErrorsOnBogusFlag(t *testing.T) {
	t.Parallel()
	args := []string{"-bogus"}
	_, err := lineCounter.NewCounter(
		lineCounter.WithOutput(io.Discard),
		lineCounter.FromArgs(args),
	)
	if err == nil {
		t.Fatal("want error on bogus flag, got nil")
	}
}

func TestBytes(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("1\n2 words\n3 this time")
	c, err := lineCounter.NewCounter(lineCounter.WithInput(inputBuf))
	if err != nil {
		t.Fatal(err)
	}
	want := 21
	got := c.Bytes()

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}

}

//func TestByteWordError(t *testing.T) {
//	t.Parallel()
//	args := []string{"-w", "-b", "testdata/three_lines.txt"}
//	_, err := lineCounter.NewCounter(
//		lineCounter.FromArgs(args),
//	)
//	if err == nil {
//		t.Fatal("want error on -w and -b, got nil")
//	}
//}
