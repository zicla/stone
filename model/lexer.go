package model

import (
	"regexp"
	"fmt"
)

type Lexer struct {
}

func (l *Lexer) ReadLine(line string) {

	//空格 \s*
	// 注释 //
	// 字符串 "(\\"|\\\\|\\n|[^"])*" 里面可以是 \" 或 \\ 或 \n 或 非"的任何字符
	// 标识符 [A-Z_a-z][A-Z_a-z0-9]*|==|<=|>=|&&|\|\| 不以数字开头 或 == 或 <= 或 >= 或 && 或 || 或
	// !"#$%&'()*+,-./:;<=>?@[\]^_`{|}~ => !"#\$%&'\(\)\*\+,-\./:;<=>\?@\[\\\]\^_`\{\|\}~
	reg, _ := regexp.Compile(`\s*(//.*)|([0-9]+)|([a-z]+)|("(\\"|\\\\|\\n|[^"])*")|([A-Z_a-z][A-Z_a-z0-9]*|==|<=|>=|&&|\|\||[\=\+\-\*\/\(\)\{\}\[\]\|\&\~\?\:\.\%\#\@\;])`)

	all := reg.FindAllString(line, -1)
	fmt.Printf("%q\n", all)

	allIndex := reg.FindAllStringIndex(line, -1);


	//判断是不是每个部分都是连贯的。


	for _, v := range allIndex {
		fmt.Println(v)
	}

	subMatch := reg.FindAllSubmatch([]byte(line), -1)
	fmt.Println("FindSubmatch", subMatch)
	for i, v := range subMatch {
		fmt.Printf("i = %d \n", i)
		for _, v1 := range v {
			fmt.Println(string(v1))
		}
	}

}
