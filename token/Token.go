package token

type Token interface {
	GetLineNumber() int
	GetNumber() int
	GetText() string
	IsIdentifier() bool
	IsNumber() bool
	IsString() bool
}