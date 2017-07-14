package model

type IdToken struct {
	Token
	Text string
}

func (i *IdToken) IsIdentifier() bool {
	return true;
}

func (i *IdToken) GetText() string {
	return i.Text
}
