package lexer

import (
	"strings"

	tok "github.com/sgolden88/lambda-calculus-interpreter/token"
)

type Lexer struct {
	code     string
	position int
}

func NewLexer(code string) Lexer {

	return Lexer{code, 0}
}
func (lexer *Lexer) NextToken() tok.Token {
nexttok:

	char := lexer.readChar()
	if char == byte(0) {

		return tok.Token{tok.EOF, "EOF"}
	}
	switch char {
	case byte('='):
		return tok.Token{tok.EQUALS, "="}

	case byte('('):
		return tok.Token{tok.LEFT_PAREN, "("}

	case byte(')'):
		return tok.Token{tok.RIGHT_PAREN, ")"}
	case byte('.'):
		return tok.Token{tok.DOT, "."}
	case byte('\n'):
	case byte(' '):
		goto nexttok //ignore whitespace and newlines
	default:
		if !alphanumeric(char) {
			return tok.Token{tok.ERROR, string(char)}
		}
		lexeme := lexer.longLexeme()
		if strings.EqualFold(lexeme, "lambda") {
			return tok.Token{tok.LAMBDA, lexeme}
		}
		return tok.Token{tok.IDENTIFIER, lexeme}
	}
	return tok.Token{}

}

//reads character then increments the position of the lexer
func (lexer *Lexer) readChar() byte {
	if lexer.position >= len(lexer.code) {
		return (0)
	}
	lexer.position++
	return (lexer.code[lexer.position-1])
}

//reads the character that is at the lexer's position without eating it
func (lexer *Lexer) peek() byte {
	if lexer.position >= len(lexer.code) {
		return 0
	}
	return (lexer.code[lexer.position])
}

//checks if a character is alphanumeric
func alphanumeric(c byte) bool {
	return isDigit(c) || isAlpha(c)
}
func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z')
}

//reads a multi-character lexeme
func (lexer *Lexer) longLexeme() string {
	lexer.position-- //start at the already eaten token
	start := lexer.position
	for ; alphanumeric(lexer.peek()); lexer.readChar() { //eats characters until an eaten character is non-alphanumeric
	}

	return (lexer.code[start:lexer.position])
}
