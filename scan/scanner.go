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

	s.tokenList = append(s.tokenList, *NewToken(Eof, "eof", nil, s.line))
	return s.tokenList
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case '(': s.addToken(LeftParen, nil)
	case ')': s.addToken(RightParen, nil)
	case '{': s.addToken(LeftBrace, nil)
	case '}': s.addToken(RightBrace, nil)
	case '+': s.addToken(Plus, nil)
	case '-': s.addToken(Minus, nil)
	case '*': s.addToken(Star, nil)
	case '/': 
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			} 
		} else if s.match('*') {
			depth := 1 // Stores pairs of /* and */
			for !s.isAtEnd() && depth != 0 {
				if s.peek() == '*' && s.peekNext() == '/' {
					s.advance()
					s.advance()
					depth--
					continue
				}

				if s.peek() == '/' && s.peekNext() == '*' {
					s.advance()
					s.advance()
					depth++
					continue
				}

				if s.peek() == '\n' {
					s.line++
				}

				s.advance()
			}
			if depth != 0 {
				panic("UNBALANCED COMMENT BLOCK")
			}
		} else {
			s.addToken(Slash, nil)
		}
	case '.': s.addToken(Dot, nil)
	case ',': s.addToken(Comma, nil) 
	case ';': s.addToken(Semicolon, nil)
	case '=':
		if s.match('=') {
			s.addToken(EqualEqual, nil)
		} else {
			s.addToken(Equal, nil)
		}
	case '!':
		if s.match('=') {
			s.addToken(BangEqual, nil)
		} else {
			s.addToken(Bang, nil)
		}
	case '<':
		if s.match('=') {
			s.addToken(LessEqual, nil)
		} else {
			s.addToken(Less, nil)
		}
	case '>':
		if s.match('=') {
			s.addToken(GreaterEqual, nil)
		} else {
			s.addToken(Greater, nil)
		} 
	// IGNORE WHITESPACE
	case ' ', '\t', '\r': break
	case '\n':
		s.line++
	default: s.addToken(Unknown, nil)
	}
}

func (s *Scanner) addToken(t TokenType, l any) {
	text := string(s.src[s.start:s.curr])
	s.tokenList = append(s.tokenList, *NewToken(t, text, l, s.line))
}

func (s *Scanner) advance() byte {
	v := s.src[s.curr]
	s.curr++
	return v
}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return 0
	}
	return s.src[s.curr]
}

func (s *Scanner) peekNext() byte {
	if s.curr+1 >= len(s.src) {
		return 0
	}
	return s.src[s.curr+1]
}

func (s *Scanner) match(exp byte) bool {
	if s.isAtEnd() { 
		return false
	}
	if s.src[s.curr] != exp { 
		return false 
	}
	s.curr++
	return true
}

func (s *Scanner) isAtEnd() bool {
	return s.curr >= len(s.src)
}
