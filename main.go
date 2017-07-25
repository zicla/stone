package main

import (
	"stone/model"
	"fmt"
)

func main() {
	lexer := model.NewLexer();

	//codes := input.GetConsoleInputText()
	codes := "m ? 12: 23; int c = a + 200;  int a = 00100; x==y; string b = \"hello   \\ \n   world 世界你好啊\"; //I'm comment";
	err := lexer.ReadLine(codes)
	if err != nil {
		fmt.Printf("词法解析时出错:%s\n", err.Error())
	}

}
