package prune

import (
	"flag"
	"fmt"
	"github.com/spf13/afero"
	"io"
	"os"
	"time"
)

type pruneConfig struct {
	age      int
	output   io.Writer
	path     string
	cfd      bool
	criteria bool
	data     []string
	fs       afero.Fs
}

type option func(*pruneConfig) error

func WithArgs(args []string) option {
	return func(p *pruneConfig) error {
		fset := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
		age := fset.Int("age", 10, "Age in Days")
		cfd := fset.Bool("cfd", false, "Candidate for deletion")

		fset.SetOutput(p.output)
		err := fset.Parse(args)
		if err != nil {
			return err
		}
		p.age = *age
		p.cfd = *cfd
		args = fset.Args()
		if len(args) < 1 {
			return nil
		}
		p.path = args[0]
		p.fs = afero.NewBasePathFs(afero.NewOsFs(), p.path)
		return nil
	}

}

func NewPruneConfig(opts ...option) (*pruneConfig, error) {
	p := pruneConfig{age: 0, output: os.Stdout, path: ""}
	for _, opt := range opts {
		err := opt(&p)
		if err != nil {
			return &pruneConfig{}, err
		}
	}
	return &p, nil
}

func (p pruneConfig) FilePrune(fs afero.Fs) {
	afero.Walk(fs, ".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		now := time.Now()
		modDate := info.ModTime()
		age := int(now.Sub(modDate).Hours() / 24)
		if age == p.age && p.criteria {
			err = fs.Remove(path)
			if err != nil {
				return err
			}
		}
		return nil
	})

}

func (p *pruneConfig) CandidateForDeletion(size int64) {
	if size > 10 {
		p.criteria = true
	}
}

func (p *pruneConfig) FileFind(fs afero.Fs) []string {
	afero.Walk(fs, ".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		now := time.Now()
		modDate := info.ModTime()
		age := int(now.Sub(modDate).Hours() / 24)
		p.CandidateForDeletion(info.Size())
		if age == p.age {
			if p.criteria {
				p.data = append(p.data, path+" candidate for deletion")
			} else {
				p.data = append(p.data, path+" Not a candidate for deletion")
			}
		}
		if age != p.age {
			p.data = append(p.data, path+" too new")
		}

		return nil
	})

	return p.data
}

func RunFs() {
	p, err := NewPruneConfig(WithArgs(os.Args[1:]))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if !p.cfd {
		p.FilePrune(p.fs)
	} else {
		data := p.FileFind(p.fs)
		for _, d := range data {
			fmt.Println(d)
		}
	}

}
