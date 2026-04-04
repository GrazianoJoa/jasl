package scan

type Scanner struct {
	src []byte
	tokenList []Token

	start int
	curr int
	line int
}

func NewScanner(src []byte) (*Scanner) {
	return &Scanner{
		src: src,
		tokenList: []Token{},
		start: 0,
		curr: 0,
		line: 1,
	}
}

func (s *Scanner) ScanTokens() []Token {
	for !s.isAtEnd() {
		s.start = s.curr
		s.scanToken()
	}

	s.tokenList = append(s.tokenList, *NewToken(Eof, "eof", s.line))
	return s.tokenList
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case '(': s.addToken(LeftParen, "(")
	case ')': s.addToken(RightParen, ")")
	case '\n': s.addToken(NewLine, "newline") 
	default: s.addToken(Unknown, "unknown")
	}
}

func (s *Scanner) addToken(t TokenType, l string) {
	// l := string(s.src[s.start:s.curr])
	s.tokenList = append(s.tokenList, *NewToken(t, l, s.line))
}

func (s *Scanner) advance() byte {
	v := s.src[s.curr]
	s.curr++
	return v
}

func (s *Scanner) isAtEnd() bool {
	return s.curr >= len(s.src)
}

