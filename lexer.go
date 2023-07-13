package envfile

const (
	ILLEGAL     = "ILLEGAL"
	EOF         = "EOF"
	IDENT       = "IDENT"
	INT         = "INT"
	LESSTHAN    = "<"
	GREATERTHAN = ">"
	EQUAL       = "="
	NOTEQUAL    = "!="
	SEMICOLON   = ";"
	// LEFTPAREN = "("
	// RIGHTPAREN = ")"
	LEFTSQUIRLY  = "{"
	RIGHTSQUIRLY = "}"
	IF           = "if"
	ELSE         = "else"
	BANG         = "!"

	FOREACH   = "foreach"
	CHANGE    = "change"
	TO        = "to"
	DELETE    = "delete"
	WITHVALUE = "with value"
	// WITH = "with"
	// VALUE = "value"
)

var keywords = map[string]string{
	"if":      IF,
	"else":    ELSE,
	"foreach": FOREACH,
	"change":  CHANGE,
	"to":      TO,
	"delete":  DELETE,
	// "with":     WITH,
	// "value":    VALUE,
	"with value": WITHVALUE,
}

func lookupIdent(ident string) string {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

type Token struct {
	TokenType string
	Literal   string
}

type Lexer struct {
	input   string
	char    byte
	pos     int
	readPos int
}

func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPos]
	}
	l.pos = l.readPos
	l.readPos += 1
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: input}
}

func (l *Lexer) Lex() *[]Token {
	tokens := []Token{}

	for range l.input {
		token := l.lexOne()
		if token.TokenType == EOF {
			break
		}

		tokens = append(tokens, token)
	}

	return &tokens
}

func (l *Lexer) lexOne() Token {
	// keep looping until we return a token
	for {
		switch l.char {
		case '\n':
			continue

		case ';':
			return Token{SEMICOLON, ";"}

		case '=':
			return Token{EQUAL, "="}

		case '<':
			return Token{LESSTHAN, "<"}

		case '>':
			return Token{GREATERTHAN, ">"}

		case '{':
			return Token{LEFTSQUIRLY, "{"}

		case '}':
			return Token{RIGHTSQUIRLY, "}"}

		case '!':
			return Token{BANG, "!"}

		default:
			if l.char == 0 {
				continue
			} else if 48 <= l.char && l.char <= 57 {
				l.backup()
				literal := l.lexInt()
				return Token{INT, literal}
			} else if isLetter(l.char) {
				l.backup()
				literal := l.lexIdent()

				// check if the identifier is a keyword
				if tok := lookupIdent(literal); tok != IDENT {
					return Token{tok, literal}
				}

				return Token{IDENT, literal}
			} else {
				return Token{ILLEGAL, string(l.char)}
			}
		}
	}
}

func (l *Lexer) backup() {
	l.readPos = l.pos
	l.pos -= 1
}

func (l *Lexer) lexInt() string {
	var literal string
	for {
		l.readChar()

		if 48 <= l.char && l.char <= 57 {
			literal += string(l.char)
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
		l.readChar()

		if isLetter(l.char) {
			literal += string(l.char)
		} else {
			// scanned something not in the identifier
			l.backup()
			return literal
		}
	}
}
