package main

import "stone/model"

func main() {
	lexer := model.Lexer{}
	lexer.ReadLine("x = m ? 12: 23; int c = a + 200; int a = 00100; x==y; string b = \"hello \\ \n \"  world\"; //I'm comment")
}
