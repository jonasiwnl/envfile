package parser

type tokenType int

const (
	tokenIllegal tokenType = iota
	tokenEOF
	tokenIdent
	tokenInt
	tokenLessThan
	tokenGreaterThan
	tokenEqual
	tokenNotEqual
	tokenSemicolon
	tokenLParen
	tokenRParen
	tokenLSquirly
	tokenRSquirly
	tokenTrue
	tokenFalse
	tokenIf
	tokenElse

	tokenField // Idk? this should just be a string

	tokenForEach
	tokenChange
	tokenTo
	tokenDelete
	tokenWith
	tokenValue
)

type token struct {
	tokenType tokenType
	text      string
}

var keywords = map[string]tokenType{
	"true":  tokenTrue,
	"false": tokenFalse,
	"if":    tokenIf,
	"else":  tokenElse,

	"foreach": tokenForEach,
	"change":  tokenChange,
	"to":      tokenTo,
	"delete":  tokenDelete,
	"with":    tokenWith,
	"value":   tokenValue,
}

const eof = -1

type Lexer struct {
	input  string
	start  int
	pos    int
	width  int
	tokens chan token
}

func NewLexer(input string) (*Lexer, chan token) {
	l := &Lexer{
		input:  input,
		tokens: make(chan token),
	}
	go l.run() // Concurrently run state machine.
	return l, l.tokens
}

// stateFn represents the state of the lexer
// as a function that returns the next state.
type stateFn func(*Lexer) stateFn

// run lexes the input by executing state functions
// until the state is nil.
func (l *Lexer) run() {
	for state := lex; state != nil; {
		state = state(l)
	}
	close(l.tokens)
}
