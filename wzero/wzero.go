package wzero

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type zeroConfig struct {
	input string
	size  int64
}

type option func(*zeroConfig) error

func NewZeroConfig(opts ...option) (*zeroConfig, error) {
	z := zeroConfig{size: 0}
	for _, opt := range opts {
		err := opt(&z)
		if err != nil {
			return &zeroConfig{}, err
		}
	}
	return &z, nil
}

func WithInput(input string) option {
	return func(z *zeroConfig) error {
		if input == "" {
			return errors.New("nil input reader")
		}
		z.input = input
		return nil
	}
}

func WithArguments(args []string) option {
	return func(z *zeroConfig) error {
		fset := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
		fsize := fset.Int64("size", 0, "Number of Zeros to write")
		fset.SetOutput(os.Stdout)
		err := fset.Parse(args)
		if err != nil {
			return err
		}
		args = fset.Args()
		if len(args) < 1 {
			return errors.New("no args provided")
		}

		z.input = args[0]
		z.size = *fsize
		return nil
	}
}
func WriteZero(path string, bytes int64) int {
	f, err := os.OpenFile(path, os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fileinfo, err := f.Stat()
	initialSize := fileinfo.Size()
	err = f.Truncate(bytes + initialSize)
	fileinfo, err = f.Stat()
	truncateSize := fileinfo.Size()
	if err != nil {
		panic(err)
	}
	return int(truncateSize)
}

func (z zeroConfig) WriteZerosFromCli() int {
	f, err := os.Create(z.input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fileinfo, err := f.Stat()
	initialSize := fileinfo.Size()
	err = f.Truncate(z.size + initialSize)
	fileinfo, err = f.Stat()
	truncateSize := fileinfo.Size()
	if err != nil {
		log.Fatal(err)
	}
	return int(truncateSize)
}

func RunZero() {
	z, err := NewZeroConfig(WithArguments(os.Args[1:]))
	if err != nil {
		log.Fatal(err)
	}
	data := z.WriteZerosFromCli()
	fmt.Println("Zeros Written " + strconv.Itoa(data))
}
