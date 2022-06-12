package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"testing/iotest"
)

type errWriter struct {
	Input io.Reader
}

func (r errWriter) Write(data []byte) (int, error) {
	return 0, errors.New("failed to write")
}

func main() {

	// A reader that always returns a custom error.
	r := iotest.ErrReader(errors.New("custom error"))
	n, err := r.Read(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("n:   %d\nerr: %q\n", n, err)

	//ew := errWriter{os.Stdin}
	//_, err := ew.Write([]byte("hello"))
	//if err != nil {
	//	log.Fatal(err)
	//}

}
