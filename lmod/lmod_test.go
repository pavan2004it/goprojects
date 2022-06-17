package lmod_test

import (
	"lmod"
	"testing"
	"testing/fstest"
	"time"
)

func TestFindModified(t *testing.T) {
	t.Parallel()
	fsys := fstest.MapFS{
		"file.go":                {[]byte(""), 0644, time.Date(2022, 5, 1, 12, 30, 0, 0, time.UTC), ""},
		"subfolder/subfolder.go": {[]byte(""), 0644, time.Date(2022, 5, 1, 12, 30, 0, 0, time.UTC), ""},
		"subfolder2/another.go":  {[]byte(""), 0644, time.Date(2022, 5, 1, 12, 30, 0, 0, time.UTC), ""},
		"subfolder2/file.go":     {[]byte(""), 0644, time.Date(2022, 6, 1, 12, 30, 0, 0, time.UTC), ""},
	}
	want := 3
	got := lmod.Files(fsys)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
