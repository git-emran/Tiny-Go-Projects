package lexer

import (
	"testing"

	"github.com/git-emran/tiny-go-projects/interpreter-in-go/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
	}
}
