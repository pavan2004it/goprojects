package main

import "fmt"

func main() {
	var title string
	var copies int
	var edition string
	var author string
	var inStock bool
	var royaltyPercentage float64
	var specialOffer bool
	var discount float64
	inStock = true
	royaltyPercentage = 12.5
	title = "For the Love of Go"
	copies = 99
	edition = "First Edition"
	author = "John"
	specialOffer = true
	discount = 10
	fmt.Println(title)
	fmt.Println(copies)
	fmt.Println(edition)
	fmt.Println(author)
	fmt.Println(inStock)
	fmt.Println(royaltyPercentage)
	fmt.Println(specialOffer)
	fmt.Println(discount)
}
