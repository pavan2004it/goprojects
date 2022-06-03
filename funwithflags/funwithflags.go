package funwithflags

import (
	"flag"
	"fmt"
	"strings"
)

type greet struct {
	word string
}

type option func(*greet) error

//func WithWord(word string) option {
//	return func(g *greet) error {
//		if word == "" {
//			return errors.New("nil input not accepted")
//		}
//		g.word = word
//		return nil
//	}
//}

func FlagParse() {
	//g := greet{word: "i am iron man"}
	capitalize := flag.Bool("c", false, "Capitalizes Strings")
	flag.Parse()
	if *capitalize {
		args := flag.Args()
		for _, arg := range args {
			fmt.Println(strings.ToUpper(arg))
		}
	} else {
		flag.PrintDefaults()
	}

}
