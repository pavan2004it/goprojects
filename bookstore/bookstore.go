package bookstore

import (
	"errors"
	"fmt"
)

// Type Catalog for mapping

type Catalog map[int]Book

// Extended catalog

type ExtendedCatalog map[int]Book

// Extended Book with additional properties

type Book struct {
	Title           string
	Author          string
	Copies          int
	ID              int
	PriceCents      int
	DiscountPercent int
	category        Category
}

type Category int

const (
	CategoryAutobiography Category = iota
	CategoryLargePrintRomance
	CategoryParticlePhysics
)

var validCategory = map[Category]bool{
	CategoryAutobiography:     true,
	CategoryLargePrintRomance: true,
	CategoryParticlePhysics:   true,
}

// Buy function updates the copies of the book
//and returns modified book

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("negative copies not supported")
	}
	b.Copies--
	return b, nil
}

// GetAllBooks gets the list of available books

func (c ExtendedCatalog) GetAllBooks() []Book {
	var result []Book
	for _, b := range c {
		result = append(result, b)
	}
	return result
}

// Func GetBook returns a book with a unique id

func (c Catalog) GetBook(id int) (Book, error) {
	b, ok := c[id]
	if !ok {
		return Book{}, fmt.Errorf("id %d doesn't exist", id)
	}
	return b, nil
}

// Func NetPriceCents returns the price of the book minus the discount

func (b Book) NetPriceCents() int {
	saving := b.DiscountPercent * b.PriceCents / 100
	return b.PriceCents - saving
}

// Func SetPriceCents sets the price for the book

func (b *Book) SetPriceCents(price int) error {
	if price < 0 {
		return fmt.Errorf("price %d cannot be negative", price)
	}
	(*b).PriceCents = price
	return nil
}

func (b Book) Category() Category {
	return b.category
}

func (b *Book) SetCategory(category Category) error {
	if !validCategory[category] {
		return fmt.Errorf("unknown category %v",
			category)
	}
	(*b).category = category
	return nil
}
