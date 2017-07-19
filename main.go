package main

import (
	"stone/model"
	"stone/input"
)

func main() {
	lexer := model.Lexer{}

	// x = m ? 12: 23; int c = a + 200\nint a = 00100; x==y; string b = \"hello \\ \n \"  world\"; //I'm comment

	codes := input.GetConsoleInputText();



	lexer.ReadLine(codes)
}
