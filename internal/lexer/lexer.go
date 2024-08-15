package lexer

import (
	"github.com/danbrakeley/hai/internal/token"
)

type Lexer struct {
	input        string
	position     int  // current reading position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.New(token.EQ, "==")
		} else {
			tok = token.New(token.ASSIGN, l.ch)
		}
	case '+':
		tok = token.New(token.PLUS, l.ch)
	case '-':
		tok = token.New(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.New(token.NOT_EQ, "!=")
		} else {
			tok = token.New(token.BANG, l.ch)
		}
	case '*':
		tok = token.New(token.ASTERISK, l.ch)
	case '/':
		tok = token.New(token.SLASH, l.ch)
	case '<':
		tok = token.New(token.LT, l.ch)
	case '>':
		tok = token.New(token.GT, l.ch)
	case ';':
		tok = token.New(token.SEMICOLON, l.ch)
	case '(':
		tok = token.New(token.LPAREN, l.ch)
	case ')':
		tok = token.New(token.RPAREN, l.ch)
	case ',':
		tok = token.New(token.COMMA, l.ch)
	case '{':
		tok = token.New(token.LBRACE, l.ch)
	case '}':
		tok = token.New(token.RBRACE, l.ch)
	case 0:
		tok = token.New(token.EOF, "")
	default:
		switch {
		case isValidStartToIdent(l.ch):
			tok = token.NewIdent(l.readIdentifier())
			// early out so we don't skip the next char
			return tok
		case isDigit(l.ch):
			lit := l.readNumber()
			if isValidStartToIdent(l.ch) {
				lit = lit + l.readIdentifier()
				tok = token.New(token.ILLEGAL, lit)
			} else {
				tok = token.New(token.INT, lit)
			}
			// early out so we don't skip the next char
			return tok
		default:
			tok = token.New(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.ch) {
		l.readChar()
	}
}

// readIdentifier assumes current char is a valid start to an identifier
func (l *Lexer) readIdentifier() string {
	position := l.position
	l.readChar()
	for isValidBodyOfIdent(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isValidStartToIdent(ch byte) bool {
	return isLetter(ch) || ch == '_'
}

func isValidBodyOfIdent(ch byte) bool {
	return isLetter(ch) || ch == '_' || isDigit(ch)
}
