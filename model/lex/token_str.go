package lex

type StrToken struct {
	Token
	Literal string
}

func NewStrToken(Literal string) *StrToken {
	return &StrToken{Literal: Literal}
}

func (s *StrToken) IsString() bool {
	return true;
}

func (s *StrToken) GetText() string {
	return s.Literal
}
