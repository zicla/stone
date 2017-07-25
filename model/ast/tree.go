package ast

//对应java的ASTree
type ITree interface {
	Child(i int) ITree;
	NumChildren() int;
	Location() string;
}

type Tree struct {
}

func (tree *Tree) Child(index int) (iTree ITree) {

	return nil
}

func (tree *Tree) NumChildren() (num int) {

	return 0
}

func (tree *Tree) Location() (str string) {

	return nil
}
