package token

//go:generate enumer -type=TokenType -json -transform=snake
type TokenType uint16

type Token struct {
	Literal string
	Type    TokenType
}

const (
	ILLEGAL TokenType = iota
	EOF

	// Identifiers + literals
	IDENT
	INT

	// Operators
	ASSIGN
	PLUS
	MINUS
	BANG
	ASTERISK
	SLASH
	LT
	GT
	EQ
	NOT_EQ

	// Delimiters
	COMMA
	SEMICOLON
	LPAREN
	RPAREN
	LBRACE
	RBRACE

	// Keywords
	FUNCTION
	LET
	TRUE
	FALSE
	IF
	ELSE
	RETURN
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
