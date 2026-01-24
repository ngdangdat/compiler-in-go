package parser

import (
	"compiler-in-go/monkey/ast"
	"compiler-in-go/monkey/lexer"
	"compiler-in-go/monkey/token"
	"fmt"
)

type Parser struct {
	l lexer.LexerItf
	curToken token.Token
	peekToken token.Token
}

func New(l lexer.LexerItf) *Parser {
	p := &Parser{l: l}
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
	// statements
	// if
	// return
	// let
	// statements.push
	program := &ast.Program{}
	for p.curToken.Type != token.EOF {
		stmt, err := p.parseStatement()
		if err != nil {
			panic("Error while parsing %v", p.curToken)
		}
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
	}

	return program
}

func (p *Parser) parseStatement() (ast.Statement, error) {
	if p.curToken.Type == token.LET {
		return p.parseLetStatement()
	}

	return nil, nil
}

func (p *Parser) parseLetStatement() (*ast.LetStatement, error) {
	ls := &ast.LetStatement{ Token: p.curToken }
	p.nextToken()

	if p.curToken.Type != token.IDENT {
	}
	identifier := parseIdentifier()
	ls.Name = identifier.(*ast.Identifier)
	p.nextToken()
	if p.curToken.Type != token.EQ {
		return nil, fmt.Errorf("invalid operator")
	}

	return ls, nil
}

func parseIdentifier() ast.Expression {

}
