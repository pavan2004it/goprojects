package main

import (
	"findmodf"
	"fmt"
	"os"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.AddDate(0, 0, -30))
	fmt.Println(lmod.Files(os.DirFS(os.Args[1])))
}
