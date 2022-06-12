package multiCounter

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type counter struct {
	wordCount bool
	input     io.Reader
	output    io.Writer
	verbose   bool
	byteCount bool
}

type option func(*counter) error

func WithInput(input io.Reader) option {
	return func(c *counter) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		c.input = input
		return nil
	}
}
func WithOutput(output io.Writer) option {
	return func(c *counter) error {
		if output == nil {
			return errors.New("nil output writer")
		}
		c.output = output
		return nil
	}
}

func FromArgs(args []string) option {
	return func(c *counter) error {
		fset := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
		wordCount := fset.Bool("w", false, "Count words instead of lines")
		verbose := fset.Bool("v", false, "Print Words or lines with more verbosity")
		byteCount := fset.Bool("b", false, "Count bytes instead of lines")
		fset.SetOutput(c.output)
		err := fset.Parse(args)
		if err != nil {
			return err
		}
		c.wordCount = *wordCount
		c.verbose = *verbose
		c.byteCount = *byteCount
		args = fset.Args()
		if len(args) < 1 {
			return nil
		}
		f, err := os.Open(args[0])
		if err != nil {
			return err
		}

		c.input = f

		return nil
	}
}

func NewCounter(opts ...option) (counter, error) {
	c := counter{input: os.Stdin, output: os.Stdout}
	for _, opt := range opts {
		err := opt(&c)
		if err != nil {
			return counter{}, err
		}
	}
	return c, nil
}

func (c counter) MultipleTypeCount() {
	input := c.input
	bytes := 0
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		bytes++
	}

	fmt.Println(bytes)

}

func (c counter) Words() int {
	words := 0
	scanner := bufio.NewScanner(c.input)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words++
	}

	return words
}

func (c counter) Bytes() int {
	bytes := 0
	scanner := bufio.NewScanner(c.input)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		bytes++
	}

	return bytes
}

func RunCLI() {
	c, err := NewCounter(
		FromArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(c.Words())
	c, err = NewCounter(
		FromArgs(os.Args[1:]),
	)
	fmt.Println(c.Bytes())
}
