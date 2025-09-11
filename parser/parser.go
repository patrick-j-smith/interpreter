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

function parseProgram() {
	program = newProgramASTNode()

	advanceTokens()

	for (currentToken() != EOF_TOKEN) {
		statement = null

		if (currentToken() == LET_TOKEN) {
			statement = parseLetStatement()
		} else if (currentToken() == RETURN_TOKEN) {
			statement = parseReturnStatement()
		} else if (currentToken() == EQUAL_TOKEN) {
			statement = parseIfStatement()
		}

		if (statement != nil) {
			program.Statements.push(statement)
		}

		advanceTokens()
	}

	return program
}

function parseLetStatement() {
	advanceTokens()

	identifier = parseIdentifier()

	advanceTokens()

	if currentToken() != EQUAL_TOKEN {
		parseError("no equal sign!")
		return null
	}

	advanceTokens()

	value = parseExpression()

	variableStatement = newVariableStatementASTNode()
	variableStatement.identifier = identifier
	variableStatement.expression = value
	return variableStatement
}

function parseIdentifier() {
	identifier = newIdentifierASTNode()
	identifier.token = currentToken()
	return identifier
}

function parseExpression() {
	if (currentToken() == INTEGER_TOKEN) {
		if (nextToken() == PLUS_TOKEN) {
			return parseOperatorExpression()
		} else if (nextToken() == SEMICOLON_TOKEN) {
			return parseIntegerLiteral()
		}
	} else if (currentToken() == LEFT_PAREN) {
		return parseGroupedExpression()
	}
}

function parseOperatorExpression() {
	operatorExpression = newOperatorExpression()

	operatorExpression.left = parseIntergerLiteral()
	advanceTokens()
	operatorExpression.operator = currentToken()
	advanceTokens()
	operatorExpression.right = parseIntegerLiteral()

	return operatorExpression
}

