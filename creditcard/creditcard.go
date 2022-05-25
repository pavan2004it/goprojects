package creditcard

import "fmt"

type card struct {
	number string
}

func New(number string) (card, error) {
	if number == "" {
		return card{}, fmt.Errorf("card number must not be empty")
	}
	return card{number}, nil
}

func (cc *card) Number() string {
	return cc.number
}
