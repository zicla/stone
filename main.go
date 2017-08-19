package main

import (
	"fmt"
	"stone/input"
	"stone/model/lex"
)

func main() {
	lexer := lex.NewLexer()

	/*
sum = 0
i = 1
while i < 10 {
        sum = sum + i
        i = i + 1
}

sum
	 */

	codes := input.GetConsoleInputText()
	//codes := "m ? 12: 23; int c = a + 200;  int a = 00100; x==y; string b = \"hello   \\ \n   world 世界你好啊\"; //I'm comment";
	err := lexer.ReadLine(codes)
	if err != nil {
		fmt.Printf("词法解析时出错:%s\n", err.Error())
	} else {
		lexer.PrintQueue()
	}

}
