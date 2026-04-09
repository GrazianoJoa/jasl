package scan

type TokenType int

const (
  LeftParen TokenType = iota
  RightParen
  LeftBrace
  RightBrace

  Plus
  Minus
  Star
  Slash

  Dot
  Comma
  Semicolon

	Equal
	EqualEqual
	Bang
	BangEqual
	Less
	LessEqual
	Greater
	GreaterEqual

	Identifier
	String
	Number

	If 
	Else 
	True
	False
	Nil

	Var
	For
	While

	Fun 
	Return
	Print

	Class
	Super
	This

	Eof
	NewLine
	Unknown
)

var tokenTypeNames = map[TokenType] string {
	LeftParen:    "LEFT_PAREN",
	RightParen:   "RIGHT_PAREN",
	LeftBrace:    "LEFT_BRACE",
	RightBrace:   "RIGHT_BRACE",

	Plus:         "PLUS",
	Minus:        "MINUS",
	Star:         "STAR",
	Slash:        "SLASH",

	Dot:          "DOT",
	Comma:        "COMMA",
	Semicolon:    "SEMICOLON",

	Equal:        "EQUAL",
	EqualEqual:   "EQUAL_EQUAL",
	Bang:         "BANG",
	BangEqual:    "BANG_EQUAL",
	Less:         "LESS",
	LessEqual:    "LESS_EQUAL",
	Greater:      "GREATER",
	GreaterEqual: "GREATER_EQUAL",

	Identifier:   "IDENTIFIER",
	String:       "STRING",
	Number:       "NUMBER",

	If:           "IF",
	Else:         "ELSE",
	True:         "TRUE",
	False:        "FALSE",
	Nil:          "NIL",

	Var:          "VAR",
	For:          "FOR",
	While:        "WHILE",

	Fun:          "FUN",
	Return:       "RETURN",
	Print:        "PRINT",

	Class:        "CLASS",
	Super:        "SUPER",
	This:         "THIS",

	Eof:          "EOF",
	NewLine:      "NEWLINE",
	Unknown:      "UNKNOWN",
}
