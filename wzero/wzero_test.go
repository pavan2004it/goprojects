package wzero_test

import (
	"testing"
	"wzero"
)

func TestWriteZero(t *testing.T) {
	input := "testdata/sample.txt"
	byteSize := int64(10000)
	want := 20000
	got := wzero.WriteZero(input, byteSize)
	if want != got {
		t.Errorf("want %d got %d", want, got)
	}
}

func TestWithArguments(t *testing.T) {
	input := []string{"-size", "10000", "testdata/new.txt"}
	w, err := wzero.NewZeroConfig(wzero.WithArguments(input))
	if err != nil {
		t.Errorf("error: %v", err)
	}
	want := 10000
	got := w.WriteZerosFromCli()
	if want != got {
		t.Errorf("want %d got %d", want, got)
	}

}

func TestWithEmptyArguments(t *testing.T) {
	input := []string{}
	_, err := wzero.NewZeroConfig(wzero.WithArguments(input))
	if err == nil {
		t.Errorf("wanted error got nil")
	}

}

func TestFromArgsErrorsOnBogusFlag(t *testing.T) {
	t.Parallel()
	args := []string{"-bogus"}
	_, err := wzero.NewZeroConfig(
		wzero.WithArguments(args),
	)
	if err == nil {
		t.Fatal("want error on bogus flag, got nil")
	}
}
