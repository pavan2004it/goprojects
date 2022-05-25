package findGo_test

import (
	"archive/zip"
	"findGo"
	"os"
	"testing"
	"testing/fstest"
)

func TestFilesOnDisk(t *testing.T) {
	t.Parallel()
	fsys := os.DirFS("testdata/findGo")
	want := 4
	got := findGo.Files(fsys)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestFilesInMemory(t *testing.T) {
	t.Parallel()
	fsys := fstest.MapFS{
		"file.go":                {},
		"subfolder/subfolder.go": {},
		"subfolder2/another.go":  {},
		"subfolder2/file.go":     {},
	}
	want := 4
	got := findGo.Files(fsys)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func BenchmarkFilesOnDisk(b *testing.B) {
	fsys := os.DirFS("testdata/findGo")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findGo.Files(fsys)
	}
}

func BenchmarkFilesInMemory(b *testing.B) {
	fsys := fstest.MapFS{
		"file.go":                {},
		"subfolder/subfolder.go": {},
		"subfolder2/another.go":  {},
		"subfolder2/file.go":     {},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findGo.Files(fsys)
	}
}

func BenchmarkZipOnDisk(b *testing.B) {
	fsys, _ := zip.OpenReader(
		"testdata/findgo.zip")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findGo.Files(fsys)
	}
}

func TestFilesInZIP(t *testing.T) {
	t.Parallel()
	fsys, err := zip.OpenReader(
		"testdata/findgo.zip")
	if err != nil {
		t.Fatal(err)
	}
	want := 4
	got := findGo.Files(fsys)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
