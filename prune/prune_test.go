package prune_test

import (
	"github.com/spf13/afero"
	"prune"
	"reflect"
	"testing"
)

func TestFilePruneWithArgs(t *testing.T) {
	t.Parallel()
	fsys := afero.NewMemMapFs()
	err := fsys.Mkdir("testdata", 0755)
	if err != nil {
		return
	}
	_, err = fsys.Create("testdata/sample.golang")
	if err != nil {
		t.Errorf("Error creating file: %v", err)
	}
	input := []string{"-age", "0", "testdata"}
	p, err := prune.NewPruneConfig(prune.WithArgs(input))
	if err != nil {
		t.Fatal(err)
	}
	p.FilePrune(fsys)
	_, fErr := fsys.Stat("sample.golang")
	if fErr == nil {
		t.Error("wanted error, got nil")
	}
}

func TestCandidateForDeletion(t *testing.T) {
	t.Parallel()
	fsys := afero.NewMemMapFs()
	err := fsys.Mkdir("testdata", 0755)
	if err != nil {
		return
	}
	f, _ := fsys.Create("testdata/sample.text")
	_, err = f.WriteString("sample text")
	if err != nil {
		return
	}
	p, err := prune.NewPruneConfig(prune.WithArgs([]string{"-cfd", "-age", "0", "testdata"}))
	if err != nil {
		t.Fatal(err)
	}

	want := []string{"testdata/sample.text candidate for deletion"}
	got := p.FileFind(fsys)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}

}
