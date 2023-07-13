package envfile_test

import (
	"testing"

	"github.com/jonasiwnl/envfile"
)

func TestLookupIdent(t *testing.T) {
}

func TestNextToken(t *testing.T) {
}

func TestNew(t *testing.T) {
	rsToken := envfile.Token{envfile.RIGHTSQUIRLY, "{"}

	if rsToken.Literal != "{" {
		t.Error("expected tokens to match")
	}
	if rsToken.TokenType != envfile.RIGHTSQUIRLY {
		t.Error("expected tokentypes to match")
	}
}

func TestString(t *testing.T) {
	lexer := envfile.NewLexer("foreach 'User' { change 's' to 'a' }")

	tokens := lexer.Lex()
	expectedTokens := []envfile.Token{
		{TokenType: envfile.FOREACH, Literal: "foreach"},
		{TokenType: envfile.IDENT, Literal: "User"},
		{TokenType: envfile.LEFTSQUIRLY, Literal: "{"},
		{TokenType: envfile.CHANGE, Literal: "change"},
		{TokenType: envfile.IDENT, Literal: "s"},
		{TokenType: envfile.TO, Literal: "to"},
		{TokenType: envfile.IDENT, Literal: "a"},
		{TokenType: envfile.RIGHTSQUIRLY, Literal: "}"},
	}
	if len(*tokens) != len(expectedTokens) {
		t.Error("expected lengths to match", len(*tokens), len(expectedTokens))
	}

	for i := range *tokens {
		if (*tokens)[i] != expectedTokens[i] {
			t.Error("expected tokens to match", (*tokens)[i], expectedTokens[i])
		}
	}
}
