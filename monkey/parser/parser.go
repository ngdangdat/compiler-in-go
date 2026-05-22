package parser

import (
	"github.com/ngdangdat/compiler-in-go/monkey/ast"
	"github.com/ngdangdat/compiler-in-go/monkey/lexer"
	"github.com/ngdangdat/compiler-in-go/monkey/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) Parser {
	p := Parser{l: l}

	// this makes sure curToken and peekToken are not null
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
	for p.curToken.Type != token.EOF {
		var stmt ast.Statement
		switch p.curToken.Type {
		case token.LET:
			stmt = p.parseLetStatement()
		}
		program.Statements = append(program.Statements, stmt)
	}
	return program
}

func (p *Parser) parseLetStatement() ast.LetStatement {
	ls := ast.LetStatement{Token: p.curToken}
	p.nextToken()
	ls.Name = p.parseIdentifier()
	if p.peekToken.Type != token.EQ {
	}
	p.nextToken()
	ls.Value = p.parseExpression()
	return ls
}

func (p *Parser) parseIdentifier() *ast.Identifier {
	identifier := &ast.Identifier{}
	return identifier
}

func (p *Parser) parseExpression() *ast.Identifier {
	ie := &ast.Identifier{}
	return ie
}
