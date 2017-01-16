package token

type IdToken struct {
	token
	text string
}

func NewIdToken(lineNumber int, id string) *IdToken {
	return &IdToken{token{lineNumber}, id}
}

func (self *IdToken) IsIdentifier() bool {
	return true
}

func (self *IdToken) GetText() string {
	return self.text
}
