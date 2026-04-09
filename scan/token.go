package scan

import "strconv"

type Token struct {
	toktype TokenType
	lexeme string
	literal any
	line int
}

func NewToken(t TokenType, lex string, lit any, line int) *Token {
	return &Token{
		toktype: t,
		literal: lit,
		lexeme: lex,
		line: line,
	}
}

func (t *Token) String() string {
	name, _ := tokenTypeNames[t.toktype]
	return name + " " + t.lexeme + " " + strconv.Itoa(t.line)
}
