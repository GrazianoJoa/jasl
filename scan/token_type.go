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
