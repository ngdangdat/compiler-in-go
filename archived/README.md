# Writing interpreter in Go

## Chapter 1: Lexer

Lexer tokenizes input from source code file to token. Token could be anything that makes sense in the context of the language.

Define token types
- Keywords: `let`, `fn`, `if`, `else`,...
- Symbol: `(`, `)`,...

```go
type Lexer struct {
    input string
    position int
    readPosition int
    ch byte
}
```
In lexer, there is a table to map character to token from input.
- input: the input (from source code file)
- position: current position
- ch: current character
- readPosition: the next character. Sometimes, lexer needs to read ahead to tokenize tokens (e.g: `==`, `!=`).
