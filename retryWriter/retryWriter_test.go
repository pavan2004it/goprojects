package retryWriter_test

import (
	"errors"
	"log"
	"os"
	"retryWriter"
	"testing"
)

func TestWriteRetry(t *testing.T) {
	ew, err := retryWriter.NewErrWriter(retryWriter.WithOutput(os.Stdout))
	if err != nil {
		log.Fatal(err)
	}
	want := errors.New("failed retries after 3 attempts")
	got, _ := retryWriter.WriteRetry(ew, 3)
	if errors.Is(want, got) {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestEmptyOutput(t *testing.T) {
	_, err := retryWriter.NewErrWriter(retryWriter.WithOutput(nil))
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}
