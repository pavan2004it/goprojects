package main

import (
	"fmt"
	"pipeline"
)

func main() {
	str, err := pipeline.FromFile("testdata/hello.txt").Column(2).String()
	if err != nil {
		return
	}
	fmt.Println(str)
}
