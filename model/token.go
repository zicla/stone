package model

import "fmt"

type Token struct {
	LineNumber int
}

var EOF Token = Token{-1}

const EOL string = "\\n"
const TestString = "Very good";


func (t *Token) GetLineNumber() int {
	return t.LineNumber;
}

func (t *Token) IsIdentifier() bool {
	return false;
}

func (t *Token) IsNumber() bool {
	return false;
}

func (t *Token) IsString() bool {
	return false;
}

func (t *Token) GetNumber() int {

	fmt.Printf("You shold override this method.")
	return 0;
}

func (t *Token) GetText() string {

	return "";
}


