package parser

import (
	"compiler-in-go/monkey/ast"
	"compiler-in-go/monkey/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `let x = 5;`
	l := lexer.New(input)
	parser := New(l)
	program := parser.ParseProgram()
	if program == nil {
		t.Error("program is nil")
		return
	}
	tests := []struct{
		expectedIdentifier string
	}{{"x"}}
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral is not 'let'. got=%q", s.TokenLiteral())
		return false
	}
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value is incorrect. want=%s, got=%s", name, letStmt.Name.Value)
		return false
	}

	return true
}
