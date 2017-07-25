package model

import "strconv"

type NumToken struct {
	Token
	Value int
}

func NewNumToken(Value int) *NumToken {
	return &NumToken{Value: Value}
}


func (n *NumToken) IsNumber() bool {
	return true;
}

func (n *NumToken) GetText() string {
	return strconv.Itoa(n.Value)
}
