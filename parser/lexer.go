package parser

type tokenType int

const (
	tokenIllegal tokenType = iota
	tokenEOF
	tokenIdent
	tokenInt
	tokenAssign
	tokenPlus
	tokenMinus
	tokenBang
	tokenAsterisk
	tokenSlash
	tokenLessThan
	tokenGreaterThan
	tokenEqual
	tokenNotEqual
	tokenComma
	tokenSemicolon
	tokenLParen
	tokenRParen
	tokenLSquirly
	tokenRSquirly
	tokenFunction
	tokenLet
	tokenTrue
	tokenFalse
	tokenIf
	tokenElse
	tokenReturn
)

type token struct {
	tokenType tokenType
	text      string
}
