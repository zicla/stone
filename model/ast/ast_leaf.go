package ast

import (
	"stone/model/lex"
	"strconv"
)

//对应java的ASTree
type ASTLeaf struct {
	Tree
	Token lex.Token
}

func NewLeaf(token lex.Token) *ASTLeaf {
	return &ASTLeaf{Token: token}
}

func (leaf *ASTLeaf) Child(index int) (iTree ITree) {

	return nil
}

func (leaf *ASTLeaf) NumChildren() (num int) {

	return 0
}

func (leaf *ASTLeaf) Location() (str string) {

	return "at line " + strconv.Itoa(leaf.Token.LineNumber)

}

