package scan

import "strconv"

type Token struct {
	toktype TokenType
	literal string
	line int
}

func NewToken(t TokenType, lit string, line int) *Token {
	return &Token{
		toktype: t,
		literal: lit,	
		line: line,
	}
}

func (t *Token) String() string {
	return strconv.Itoa(int(t.toktype)) + " " + t.literal + " " + strconv.Itoa(t.line); 
}
