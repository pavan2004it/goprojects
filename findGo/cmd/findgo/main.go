package main

import (
	"findGo"
	"fmt"
	"os"
)

func main() {
	fmt.Println(findGo.Files(os.DirFS(os.Args[1])))
}
