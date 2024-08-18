package parser

import (
	"fmt"

	"github.com/danbrakeley/hai/internal/ast"
	"github.com/danbrakeley/hai/internal/lexer"
	"github.com/danbrakeley/hai/internal/token"
)

type Parser struct {
	lex       *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		lex:    l,
		errors: []string{},
	}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type().String())
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lex.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = make([]ast.Statement, 0, 64)

	for !p.curToken.Is(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type() {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

// parseLetStatement assumes curToken is LET
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}
	p.nextToken()

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal()}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	p.nextToken()

	// TODO: parse expressions
	p.skipTo(token.SEMICOLON)
	if !p.expectPeek(token.SEMICOLON) {
		return nil
	}

	return stmt
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekToken.Is(t) {
		return true
	}
	p.peekError(t)
	return false
}

func (p *Parser) skipTo(t token.TokenType) {
	for !p.peekToken.Is(t) && !p.peekToken.Is(token.EOF) {
		p.nextToken()
	}
}

// parseReturnStatement assumes curToken is RETURN
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	// TODO: parse expressions
	p.skipTo(token.SEMICOLON)
	if !p.expectPeek(token.SEMICOLON) {
		return nil
	}

	return stmt
}
