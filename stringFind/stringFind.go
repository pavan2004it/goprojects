package stringFind

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type finder struct {
	input     string
	search    string
	sensitive bool
}

type option func(*finder) error

func WithInput(input string) option {
	return func(f *finder) error {

		_, err := os.Stat(input)
		if err != nil {
			return errors.New("file doesn't exist")
		}
		f.input = input
		return nil
	}
}

func WithSearch(search string) option {
	return func(f *finder) error {
		if len(search) == 0 {
			return errors.New("search string can't empty")
		}
		f.search = search
		return nil
	}
}

func WithSensitive(sensitive bool) option {
	return func(f *finder) error {
		f.sensitive = sensitive
		return nil
	}
}

func NewFinder(opts ...option) (finder, error) {
	f := finder{input: "sample.txt", search: "hello", sensitive: true}
	for _, opt := range opts {
		err := opt(&f)
		if err != nil {
			return finder{}, err
		}
	}
	return f, nil
}

func (f finder) Lines() string {
	var foundLines string
	inFile, err := os.Open(f.input)
	if err != nil {
		log.Fatal(err)
	}
	defer func(inFile *os.File) {
		Closerr := inFile.Close()
		if Closerr != nil {
			log.Fatal(Closerr)
		}
	}(inFile)

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), f.search) {
			foundLines += scanner.Text() + "\n"
		}
	}
	return foundLines
}

func (f finder) InSensitiveLines() string {
	var foundLines string
	inFile, err := os.Open(f.input)
	if err != nil {
		log.Fatal(err)
	}
	defer func(inFile *os.File) {
		Closerr := inFile.Close()
		if Closerr != nil {
			log.Fatal(Closerr)
		}
	}(inFile)

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		if strings.Contains(strings.ToLower(scanner.Text()), strings.ToLower(f.search)) {
			foundLines += scanner.Text() + "\n"
		}
	}
	return foundLines
}

func Lines() {
	f, err := NewFinder()
	if err != nil {
		panic("internal error")
	}
	if f.sensitive {
		lines := f.Lines()
		fmt.Println(lines)
	} else {
		lines := f.InSensitiveLines()
		fmt.Println(lines)
	}

}
