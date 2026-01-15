package parser

import (
	"interpreter/ast"
	"interpreter/lexer"
	"interpreter/token"
)

type Parser struct {
	l *lexer.Lexer

	currentToken token.Token
	peekToken    token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	return nil
}

func (p *Parser) parseLetStatement() *ast.Program {
	return nil
}

func (p *Parser) parseIdentifier() *ast.Program {
	return nil
}

func (p *Parser) parseExpression() *ast.Program {
	return nil
}

func (p *Parser) parseOperatorExpression() *ast.Program {
	return nil
}
