package parser

import (
	"testing"

	"github.com/ngdangdat/compiler-in-go/monkey/ast"
	"github.com/ngdangdat/compiler-in-go/monkey/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returns nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("ParseProgram() invalid number of parsed statements, expected=%d, got=%d", 3, len(program.Statements))
	}
	testParserErrors(t, p)
	tests := []struct{ expectedIdentifier string }{
		{expectedIdentifier: "x"},
		{expectedIdentifier: "y"},
		{expectedIdentifier: "foobar"},
	}
	for ti, tc := range tests {
		stmt := program.Statements[ti]
		res := testLetStatement(t, stmt, tc.expectedIdentifier)
		if !res {
			t.Fatalf("case %d failed", ti)
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Fatalf("testLetStatement expected let, got=%s", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Fatalf("testLetStatement invalid statement parsed")
		return false
	}

	if letStmt.Name.Value != name {
		t.Fatalf("testLetStatement Name.Value is invalid, expected=%s got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Fatalf("testLetStatement Name.TokenLiteral() is invalid, expected=%s got=%s", name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func testParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %s", msg)
	}
	t.FailNow()
}
