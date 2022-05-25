package main

import "fmt"

func main() {
	var x int = 12

	fmt.Println(double(x))
}

func double(x int) int {
	x *= 2
	return x
}
