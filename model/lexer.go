package model

import (
	"regexp"
	"fmt"
	"errors"
)

type Lexer struct {
	queue []IToken
}

func NewLexer() *Lexer {
	return &Lexer{queue: make([]IToken, 0, 10)}
}

func (lexer Lexer) ReadLine(line string) (err error) {

	lexer.queue = append(lexer.queue, IdToken{})

	//空格 \s*
	// 注释 //
	// 字符串 "(\\"|\\\\|\\n|[^"])*" 里面可以是 \" 或 \\ 或 \n 或 非"的任何字符
	// 标识符 [A-Z_a-z][A-Z_a-z0-9]*|==|<=|>=|&&|\|\| 不以数字开头 或 == 或 <= 或 >= 或 && 或 || 或
	// !"#$%&'()*+,-./:;<=>?@[\]^_`{|}~ => !"#\$%&'\(\)\*\+,-\./:;<=>\?@\[\\\]\^_`\{\|\}~
	reg, _ := regexp.Compile(`\s*(//.*)|([0-9]+)|("(\\"|\\\\|\\n|[^"])*")|([A-Z_a-z][A-Z_a-z0-9]*|==|<=|>=|&&|\|\||[\=\+\-\*\/\(\)\{\}\[\]\|\&\~\?\:\.\%\#\@\;])`)

	//从 pos 到 lineLength 一直搜索.
	lineLength := len(line);
	startIndex := 0;
	fmt.Printf("start to search: length = %d\n", lineLength)
	for startIndex < lineLength {

		fmt.Printf("startIndex => %d\n", startIndex);
		searchStr := line[startIndex:];
		subMatch := reg.FindSubmatch([]byte(searchStr))

		if subMatchIndex := reg.FindStringIndex(searchStr); subMatchIndex == nil {
			break
		} else {

			//判断从前面开始没有被匹配到的部分是不是空格，不是得话就报错.
			unMatchPart := searchStr[0:subMatchIndex[0]];
			if isSpace, _ := regexp.Match(`^\s*$`, []byte(unMatchPart)); !isSpace {
				err = errors.New(unMatchPart + "语法错误")
				return err
			}

			fmt.Printf("match start: %d end: %d\n", startIndex+subMatchIndex[0], startIndex+subMatchIndex[1])
			startIndex += subMatchIndex[1]
		}

		for k1, v1 := range subMatch {

			if len(v1) > 0 {

				switch k1 {
				case 1:
					fmt.Println("注释：")
					fmt.Println(string(v1))
				case 2:
					fmt.Println("数字：")
					fmt.Println(string(v1))
				case 3:
					fmt.Println("字符串：")
					fmt.Println(string(v1))
				case 4:
					fmt.Println("未知的内容：")
					fmt.Println(string(v1))
				case 5:
					fmt.Println("标识符：")
					fmt.Println(string(v1))
				}

			}

		}

		fmt.Println()
	}

	return nil
}
