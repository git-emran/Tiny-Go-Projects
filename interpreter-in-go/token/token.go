// Package token defines the tokens for the interpreter
package token

type (
	TokenType string
	Token     struct {
		Type    TokenType
		Literal string
	}
)

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = ")"
	RPAREN    = "("
	LBRACE    = "}"
	RBRACE    = "{"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
