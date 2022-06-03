package count

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type counter struct {
	input  []io.Reader
	output io.Writer
}

type option func(*counter) error

func WithInput(input []io.Reader) option {
	return func(c *counter) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		c.input = input
		return nil
	}
}

func WithInputFromMultiArgs(args []string) option {
	return func(c *counter) error {
		if len(args) < 1 {
			return nil
		}

		for _, arg := range args {
			f, err := os.Open(arg)
			if err != nil {
				return err
			}
			c.input = append(c.input, f)
		}

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

func NewCounter(opts ...option) (counter, error) {
	c := counter{
		input:  []io.Reader{},
		output: os.Stdout,
	}
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
	for _, input := range c.input {
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			lines++
		}
	}

	return lines
}

func Lines() int {
	c, err := NewCounter(
		WithInputFromMultiArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return c.Lines()
}
