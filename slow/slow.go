package slow

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type slowdata struct {
	input []io.Reader
}

type option func(*slowdata) error

func WithInput(input []io.Reader) option {
	return func(s *slowdata) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		s.input = input
		return nil
	}
}

func NewSlowData(opts ...option) (slowdata, error) {
	s := slowdata{input: []io.Reader{}}
	for _, opt := range opts {
		err := opt(&s)
		if err != nil {
			return slowdata{}, err
		}
	}
	return s, nil
}

func (s slowdata) PrintSlow() error {
	for _, input := range s.input {
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			time.Sleep(3 * time.Second)
			scanner.Text()
		}
	}
	return nil

}

func WithArgs(args []string) option {
	return func(s *slowdata) error {
		if len(args) == 0 {
			return errors.New("no args")
		}
		f, err := os.Open(args[0])
		if err != nil {
			for _, arg := range args {
				s.input = append(s.input, strings.NewReader(arg))
			}
			return nil
		}
		s.input = append(s.input, f)
		return nil
	}
}

func (s slowdata) PrintSlowWithArgs() (error, string) {
	for _, input := range s.input {
		scanner := bufio.NewScanner(input)
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			time.Sleep(3 * time.Second)
			return nil, scanner.Text()
		}
	}
	return nil, ""
}

func RunCli() {
	s, err := NewSlowData(WithArgs(os.Args[1:]))
	if err != nil {
		panic(err)
	}
	for _, input := range s.input {
		scanner := bufio.NewScanner(input)
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			time.Sleep(3 * time.Second)
			fmt.Println(scanner.Text())
		}
	}
}
