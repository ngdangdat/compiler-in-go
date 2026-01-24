package token

type TokenType string

// token types
const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	IDENT = "IDENT"
	INT = "INT"

	ASSIGN = "="
	PLUS = "+"
	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"

	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET = "LET"
)

type Token struct {
	Type TokenType
	Literal string
}

var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"let": LET,
}

// Decide if it's identifier or keyword
func LookupIdent(ident string) TokenType {
	if tokType, ok := keywords[ident]; ok {
		return tokType
	}
	return IDENT
}
