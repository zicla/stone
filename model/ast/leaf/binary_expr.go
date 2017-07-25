package leaf

import "stone/model/ast"

type BinaryExpr struct {
	ast.ASTList
}

func (binaryExpr *BinaryExpr) Operator() (str string) {
	return ast.ASTLeaf(binaryExpr.Child(1)).Token.GetText();
}

func (binaryExpr *BinaryExpr) Left() (iTree ast.ITree) {
	return binaryExpr.Child(0)
}

func (binaryExpr *BinaryExpr) Right() (iTree ast.ITree) {
	return binaryExpr.Child(2)
}
