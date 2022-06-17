package lmod

import (
	"io/fs"
	"time"
)

func Files(fsys fs.FS) (count int) {
	fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		x, _ := d.Info()
		if !x.IsDir() && x.ModTime().Before(time.Now().AddDate(0, 0, -29)) {
			count++
		}
		return nil
	})
	return count
}
