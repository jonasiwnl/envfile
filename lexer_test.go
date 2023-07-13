package streamfile_test

import (
	"testing"

	"github.com/jonasiwnl/streamfile"
)

func TestLookupIdent(t *testing.T) {
}

func TestNextToken(t *testing.T) {
}

func TestNew(t *testing.T) {
	rsToken := streamfile.Token{streamfile.RIGHTSQUIRLY, "{"}

	if rsToken.Literal != "{" {
		t.Error("expected tokens to match")
	}
	if rsToken.TokenType != streamfile.RIGHTSQUIRLY {
		t.Error("expected tokentypes to match")
	}
}

func TestString(t *testing.T) {
	lexer := streamfile.NewLexer("foreach 'User' { change 's' to 'a' }")

	tokens := lexer.Lex()
	expectedTokens := []streamfile.Token{
		{TokenType: streamfile.FOREACH, Literal: "foreach"},
		{TokenType: streamfile.IDENT, Literal: "User"},
		{TokenType: streamfile.LEFTSQUIRLY, Literal: "{"},
		{TokenType: streamfile.CHANGE, Literal: "change"},
		{TokenType: streamfile.IDENT, Literal: "s"},
		{TokenType: streamfile.TO, Literal: "to"},
		{TokenType: streamfile.IDENT, Literal: "a"},
		{TokenType: streamfile.RIGHTSQUIRLY, Literal: "}"},
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
