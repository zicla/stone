package leaf

import "stone/model/ast"

type Name struct {
	ast.ASTLeaf
}

func (name *Name) Name() (str string) {
	return name.Token.GetText()
}
