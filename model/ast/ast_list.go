package ast

type ASTList struct {
	Tree
	Children []ITree
}

func (list *ASTList) Child(index int) (iTree ITree) {

	return list.Children[index]
}

func (list *ASTList) NumChildren() (num int) {

	return len(list.Children)
}

func (list *ASTList) Location() (str string) {

	return nil
}
