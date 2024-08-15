package token

type Token struct {
	lit string
	typ TokenType
}

func New[S byte | rune | string](tokenType TokenType, s S) Token {
	return Token{typ: tokenType, lit: string(s)}
}

func NewIdent(ident string) Token {
	return Token{typ: IdentType(ident), lit: ident}
}

func (t Token) Literal() string {
	return t.lit
}

func (t Token) Type() TokenType {
	return t.typ
}

func (t Token) Is(typ TokenType) bool {
	return t.typ == typ
}

//go:generate enumer -type=TokenType -json -transform=snake
type TokenType uint8

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

func IdentType(ident string) TokenType {
	switch ident {
	case "fn":
		return FUNCTION
	case "let":
		return LET
	case "true":
		return TRUE
	case "false":
		return FALSE
	case "if":
		return IF
	case "else":
		return ELSE
	case "return":
		return RETURN
	}
	return IDENT
}
