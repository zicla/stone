package leaf

import "stone/model/ast"

type NumberLiberal struct {
	ast.ASTLeaf
}

func (numberLiberal *NumberLiberal) Value() (value int) {
	return numberLiberal.Token.GetNumber()
}
