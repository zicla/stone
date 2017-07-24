package main

import (
	"stone/model"
)

func main() {
	lexer := model.Lexer{}

	// x   = m ? 12: 23; int c = a + 200;  int a = 00100; x==y; string b = \"hello \\ \n \"  world\"; //I'm comment

	//codes := input.GetConsoleInputText();
	codes := "m ? 12: 23; int c = a + 200;  int a = 00100; x==y; string b = \"hello \\ \n \"  world 世界你好啊\"; //I'm comment";
	lexer.ReadLine(codes)
}
