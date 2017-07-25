package model

type IdToken struct {
	Token
	Text string
}

func NewIdToken(Text string) *IdToken {
	return &IdToken{Text: Text}
}
func (i *IdToken) IsIdentifier() bool {
	return true;
}

func (i *IdToken) GetText() string {
	return i.Text
}
