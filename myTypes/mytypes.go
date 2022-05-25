package myTypes

import "strings"

type MyInt int
type MyString string

// Struct with anonymous field strings.builder

type MyBuilder struct {
	strings.Builder
}

type StringUppercaser struct {
	Contents strings.Builder
}

//type MyBuilder strings.Builder

// Twice multiplies its receiver by 2 and returns
// the result.
func (i MyInt) Twice() MyInt {
	return i * 2
}

func (s MyString) MyStringLen() int {
	return len(s)
}

func (mb MyBuilder) Hello() string {
	return "Hello, Gophers!"
}

func (su StringUppercaser) ToUpper() string {
	return strings.ToUpper(su.Contents.String())
}

func (input *MyInt) Double() {
	*input *= 2
}
