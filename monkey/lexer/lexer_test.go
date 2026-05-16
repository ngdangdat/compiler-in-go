package lexer

import (
	"testing"

	"github.com/ngdangdat/compiler-in-go/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []struct {
		expectedToken   token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedToken {
			t.Fatalf("tests[%d] - tokenType wrong, expected=%q, got=%q", i, tt.expectedToken, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - tokenType wrong, expected=%q, got=%q", i, tt.expectedToken, tok.Type)
		}

	}
}
