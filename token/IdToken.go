package token

type IdToken struct {
	BaseToken
	text string
}

func NewIdToken(lineNumber int, id string) *StrToken {
	return &StrToken{BaseToken{lineNumber}, id}
}

func (self *IdToken) IsIdentifier() bool {
	return true
}

func (self *IdToken) GetText() string {
	return self.text
}
