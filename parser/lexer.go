package parser

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

type tokenType int

const (
	ILLEGAL tokenType = iota
	EOF
	IDENT
	INT
	LESSTHAN
	GREATERTHAN
	EQUAL
	NOTEQUAL
	SEMICOLON
	// LEFTPAREN
	// RIGHTPAREN
	LEFTSQUIRLY
	RIGHTSQUIRLY
	IF
	ELSE
	BANG

	FOREACH
	CHANGE
	TO
	DELETE
	WITHVALUE
	// WITH
	// VALUE
)

type token struct {
	tokenType tokenType
	text      string
}

func (t token) String() string {
	return tokenStrings[t.tokenType]
}

var tokenStrings = []string{
	ILLEGAL:      "ILLEGAL",
	EOF:          "EOF",
	IDENT:        "IDENT",
	INT:          "INT",
	LESSTHAN:     "<",
	GREATERTHAN:  ">",
	EQUAL:        "=",
	NOTEQUAL:     "!=",
	SEMICOLON:    ";",
	LEFTSQUIRLY:  "{",
	RIGHTSQUIRLY: "}",
	IF:           "if",
	ELSE:         "else",
	BANG:         "!",

	FOREACH:   "foreach",
	CHANGE:    "change",
	TO:        "to",
	DELETE:    "delete",
	WITHVALUE: "with value",
	// WITH: "with",
	// VALUE: "value",
}

type Lexer struct {
	line   int
	column int
	reader *bufio.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		line:   1,
		column: 0,
		reader: bufio.NewReader(reader),
	}
}

func (l *Lexer) Lex() {
	for {
		line, column, token := l.lexOne()
		if token.tokenType == EOF {
			break
		}

		os.Stdout.WriteString(
			fmt.Sprintf(
				"%d:%d\t%s\t%s\n",
				line,
				column,
				token.String(),
				token.text,
			))
	}
}

func (l *Lexer) lexOne() (int, int, token) {
	// keep looping until we return a token
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.line, l.column, token{EOF, ""}
			}

			log.Fatal(err)
		}

		l.column++

		switch r {
		case '\n':
			l.line++
			l.column = 0

		case ';':
			return l.line, l.column, token{SEMICOLON, ";"}

		case '=':
			return l.line, l.column, token{EQUAL, "="}

		case '<':
			return l.line, l.column, token{LESSTHAN, "<"}

		case '>':
			return l.line, l.column, token{GREATERTHAN, ">"}

		case '{':
			return l.line, l.column, token{LEFTSQUIRLY, "{"}

		case '}':
			return l.line, l.column, token{RIGHTSQUIRLY, "}"}

		case '!':
			return l.line, l.column, token{BANG, "!"}

		default:
			if unicode.IsSpace(r) {
				continue
			} else if unicode.IsDigit(r) {
				startLine, startCol := l.line, l.column
				l.backup()
				literal := l.lexInt()
				return startLine, startCol, token{INT, literal}
			} else if unicode.IsLetter(r) {
				startLine, startCol := l.line, l.column
				l.backup()
				literal := l.lexIdent()
				return startLine, startCol, token{IDENT, literal}
			} else {
				return l.line, l.column, token{ILLEGAL, string(r)}
			}
		}
	}
}

func (l *Lexer) backup() {
	if err := l.reader.UnreadRune(); err != nil {
		log.Fatal(err)
	}

	l.column--
}

func (l *Lexer) lexInt() string {
	var literal string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				// at the end of the int
				return literal
			}
		}

		l.column++
		if unicode.IsDigit(r) {
			literal += string(r)
		} else {
			// scanned something not in the integer
			l.backup()
			return literal
		}
	}
}

func (l *Lexer) lexIdent() string {
	var literal string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				// at the end of the identifier
				return literal
			}
		}

		l.column++
		if unicode.IsLetter(r) {
			literal += string(r)
		} else {
			// scanned something not in the identifier
			l.backup()
			return literal
		}
	}
}
