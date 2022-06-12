package retryWriter

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

type ErrWriter struct {
	output io.Writer
}

func (e ErrWriter) Write(data []byte) (int, error) {
	return 0, errors.New("failed to write")
}

type option func(writer *ErrWriter) error

func WithOutput(output io.Writer) option {
	return func(e *ErrWriter) error {
		if output == nil {
			return errors.New("nil output writer")
		}
		e.output = output
		return nil
	}
}

func NewErrWriter(opts ...option) (ErrWriter, error) {
	e := ErrWriter{output: os.Stdout}
	for _, opt := range opts {
		err := opt(&e)
		if err != nil {
			return ErrWriter{}, err
		}
	}
	return e, nil
}

func WriteRetry(output io.Writer, retry int) (error, string) {
	retry = 3
	bytesWritten, err := output.Write([]byte("Test Data"))
	if err != nil {
		for i := 0; i < retry; i++ {
			bytesWritten, err = output.Write([]byte("Test Data"))
			if err == nil {
				return nil, "Successfully wrote " + fmt.Sprintf("%d", bytesWritten) + " bytes"
			}
		}
		return errors.New("failed retries after " + fmt.Sprintf("%d", retry) + " attempts"), "\nFailed to write Data to output"
	}
	return nil, "Successfully wrote " + fmt.Sprintf("%d", bytesWritten) + " bytes"
}

func RunCli(retry int) {
	f, err := os.OpenFile("testdata/sample.txt", os.O_WRONLY, 0644)
	err, message := WriteRetry(f, retry)
	if err != nil {
		log.Fatal(err, message)
	}
	fmt.Println(message)
}
