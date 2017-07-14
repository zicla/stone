package main

import "stone/model"

func main() {
	lexer := model.Lexer{}
	lexer.ReadLine("inta=100;")
}
