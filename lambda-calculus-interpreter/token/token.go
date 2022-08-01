package token

type Token struct {
	T      TokenType
	Lexeme string
}

func (t *Token) String() {
	println("<Token of type ", t.T, "Literal", t.Lexeme+">")
}

type TokenType uint8

const (
	LAMBDA TokenType = iota
	EQUALS
	LEFT_PAREN
	RIGHT_PAREN
	DOT
	IDENTIFIER
	EOF
	ERROR
)
