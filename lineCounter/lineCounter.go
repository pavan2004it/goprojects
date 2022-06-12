package lineCounter

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

type counter struct {
	input     io.Reader
	output    io.Writer
	verbose   bool
	wordCount bool
	byteCount bool
	lineCount bool
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
		lineCount := fset.Bool("l", false, "Count lines instead of lines")
		fset.SetOutput(c.output)
		err := fset.Parse(args)
		if err != nil {
			return err
		}
		c.wordCount = *wordCount
		c.verbose = *verbose
		c.byteCount = *byteCount
		c.lineCount = *lineCount
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
func (c counter) Lines() int {
	lines := 0
	scanner := bufio.NewScanner(c.input)
	for scanner.Scan() {
		lines++
	}
	return lines

}
func Lines() int {
	c, err := NewCounter(FromArgs(os.Args[1:]))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return c.Lines()
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

func Words() int {
	c, err := NewCounter(FromArgs(os.Args[1:]))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return c.Words()
}

func RunCLI() {
	c, err := NewCounter(FromArgs(os.Args[1:]))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {

	case c.wordCount && c.byteCount:
		fmt.Println(strconv.Itoa(c.Words()) + " words")
		c, _ = NewCounter(FromArgs(os.Args[1:]))
		fmt.Println(strconv.Itoa(c.Bytes()) + " bytes")

	case c.wordCount && c.lineCount:
		fmt.Println(strconv.Itoa(c.Words()) + " words")
		c, _ = NewCounter(FromArgs(os.Args[1:]))
		fmt.Println(strconv.Itoa(c.Lines()) + " lines")

	case c.byteCount && c.lineCount:
		fmt.Println(strconv.Itoa(c.Bytes()) + " bytes")
		c, _ = NewCounter(FromArgs(os.Args[1:]))
		fmt.Println(strconv.Itoa(c.Lines()) + " lines")

	case c.wordCount && c.verbose:
		c.Words()
		fmt.Println(strconv.Itoa(c.Words()) + " words")

	case c.verbose:
		fmt.Println(strconv.Itoa(c.Lines()) + " lines")

	case c.wordCount:
		fmt.Println(c.Words())

	case c.byteCount:
		fmt.Println(c.Bytes())

	default:
		fmt.Println(c.Lines())
	}

}
