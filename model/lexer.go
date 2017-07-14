package model

import (
	"regexp"
	"fmt"
)

type Lexer struct {
}

func (l *Lexer) ReadLine(line string) {

	reg, _ := regexp.Compile(`([0-9]+)|([a-z]+)`)

	all := reg.FindAllString(line, 1)
	fmt.Printf("%q\n", all)

	allIndex := reg.FindAllStringIndex(line, -1);
	fmt.Printf("%q\n", allIndex)

}
