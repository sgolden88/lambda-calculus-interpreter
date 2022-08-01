package parser

import (
	"fmt"

	"github.com/sgolden88/lambda-calculus-interpreter/ast"
	"github.com/sgolden88/lambda-calculus-interpreter/lexer"
	"github.com/sgolden88/lambda-calculus-interpreter/token"
)

type Parser struct {
	lexer *lexer.Lexer
	curr  token.Token
	next  token.Token
}

func NewParser(l *lexer.Lexer) Parser {
	return Parser{l, l.NextToken(), l.NextToken()}
}

func (parser *Parser) Parse() (ast.AstNode, error) {
	//it is either a binding or an expression
	if parser.curr.T == token.EOF {
		return nil, nil
	} else if (parser.curr.T == token.IDENTIFIER) && (parser.next.T == token.EQUALS) {
		//handle binding
		identifier := ast.Identifier(parser.curr.Lexeme)
		parser.nextTok() //progresses the parser
		parser.nextTok()
		value, err := parser.parseExp()
		if err != nil {
			return nil, err
		}
		return ast.Binding{identifier, value}, nil
	} else {
		//handle expression sequence
		return parser.parseSequence()

	}
}
func (parser *Parser) parseExp() (ast.Expression, error) {

	switch parser.curr.T {
	case token.LEFT_PAREN:
		parser.nextTok() //eat the left brace
		if parser.curr.T == token.EOF {
			return nil, fmt.Errorf("Expected Right Parentheses")
		}
		ret, err := parser.parseExp()
		if err != nil {
			return nil, err
		}
		for parser.curr.T != token.RIGHT_PAREN { //No EOF token
			nextexp, err := parser.parseExp()
			if err != nil {
				return nil, err
			}
			ret = ast.Application{ret, nextexp} //progress the parser after the expression has been parsed
		}
		parser.nextTok() //eat right bracket
		return ret, nil
	case token.LAMBDA:
		if parser.next.T != token.IDENTIFIER {
			return nil, fmt.Errorf("Expected identifier, instead received: " + parser.next.Lexeme)
		}
		parser.nextTok()
		if parser.next.T != token.DOT {
			return nil, fmt.Errorf("Expected dot to seperate identifier and expression")
		}
		identifer := ast.Identifier(parser.curr.Lexeme)
		parser.nextTok()
		parser.nextTok()
		expr, err := parser.parseSequence()
		return ast.Abstraction{identifer, expr}, err
	case token.IDENTIFIER:
		ret := ast.Identifier(parser.curr.Lexeme)
		parser.nextTok() //progress the parser after expression has been parsed,sort of brokekn
		return ret, nil
	default:
		return nil, fmt.Errorf("Unexpected token: " + parser.curr.Lexeme)
	}
}
func (parser *Parser) parseSequence() (ast.Expression, error) {
	if parser.curr.T == token.EOF {
		return nil, nil
	}
	ret, err := parser.parseExp()
	if err != nil {
		return nil, err
	}
	for (parser.curr.T != token.EOF) && (parser.curr.T != token.RIGHT_PAREN) { //No EOF token
		nextexp, err := parser.parseExp()
		if err != nil {
			return nil, err
		}
		ret = ast.Application{ret, nextexp} //progress the parser after the expression has been parsed
	}
	return ret, nil

}
func (parser *Parser) nextTok() {
	parser.curr = parser.next
	parser.next = parser.lexer.NextToken()
}
