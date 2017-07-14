package model

type StrToken struct {
	Token
	Literal string
}

func (s *StrToken) IsString() bool {
	return true;
}

func (s *StrToken) GetText() string {
	return s.Literal
}
