package parser

import (
	"compiler-in-go/monkey/ast"
	"compiler-in-go/monkey/lexer"
	"compiler-in-go/monkey/token"
	"fmt"
)

type Parser struct {
	l lexer.LexerItf
	errors []string
	curToken token.Token
	peekToken token.Token
}

func New(l lexer.LexerItf) *Parser {
	p := &Parser{
		errors: []string{},
		l: l,
	}
	// jump 2 times to make sure both cur and peek tokens are not nil
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	}

	return nil
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	ls := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	ls.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	if !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return ls
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	rs := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()
	
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return rs
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if !p.peekTokenIs(t) {
		p.peekError(t)
		return false
	}
	p.nextToken()
	return true
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("invalid next token want=%s, got=%s", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
