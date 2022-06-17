package main

import (
	"fmt"
	"github.com/spf13/afero"
	"log"
	"os"
	"path/filepath"
	"time"
)

//type errWriter struct {
//	Input io.Reader
//}
//
//func (r errWriter) Write(data []byte) (int, error) {
//	return 0, errors.New("failed to write")
//}

func main() {

	// A reader that always returns a custom error.
	//r := iotest.ErrReader(errors.New("custom error"))
	//n, err := r.Read(nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("n:   %d\nerr: %q\n", n, err)

	//ew := errWriter{os.Stdin}
	//_, err := ew.Write([]byte("hello"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//ShowOldestFile()
	fstest()

}

func ShowOldestFile() {
	filepath.Walk("/Users/pavankumar/Library/Caches/go-build/00", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {

			_, _, day := info.ModTime().Date()
			daysPassed := time.Now().Day() - day
			fmt.Printf("Days passed %d: since the file %q was created since today\n", daysPassed, info.Name())

		}
		return nil
	})
}

func fstest() {
	//var AppFs = afero.NewMemMapFs()
	var AppOsFs = afero.NewOsFs()
	//err := AppFs.Mkdir("/Users/pavankumar/goprojects/sandbox/testdata/cmd", 0777)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//Oserr := AppOsFs.Mkdir("/Users/pavankumar/goprojects/sandbox/testdata/cmd", 0777)
	//if Oserr != nil {
	//	log.Fatal(Oserr)
	//}
	f, fileerr := AppOsFs.Create("/Users/pavankumar/goprojects/sandbox/testdata/cmd/test.txt")
	if fileerr != nil {
		log.Fatal(fileerr)
	}
	info, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}

	//err := AppOsFs.Remove("/Users/pavankumar/goprojects/sandbox/testdata/cmd/test.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}

	fmt.Println(info.ModTime())
}
