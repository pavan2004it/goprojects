package main

import (
	"dataMarshall"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"log"
)

func main() {
	path := "/Users/pavankumar/goprojects/dataMarshall/store.bin"
	output := dataMarshall.SetPath(path)
	want := []int{2, 3, 5, 7, 11}
	err := output.Encode(want)
	if err != nil {
		log.Fatal(err)
	}
	output.Close()
	input := dataMarshall.SetPath(path)
	var got []int
	err = input.Decode(&got)
	if err != nil {
		log.Fatal(err)
	}
	input.Close()
	if !cmp.Equal(want, got) {
		log.Fatal(cmp.Diff(want, got))
	}
	fmt.Println(want, got)
}
