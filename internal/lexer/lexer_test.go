package lexer

import (
	"fmt"
	"testing"

	"github.com/danbrakeley/hai/internal/token"
)

func TestNextToken_IndividualTokens(t *testing.T) {
	var cases = []struct {
		token           string
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{"#", token.ILLEGAL, "#"},
		{"", token.EOF, ""},
		{"a", token.IDENT, "a"},
		{"aBcD", token.IDENT, "aBcD"},
		{"ab_YZ", token.IDENT, "ab_YZ"},
		{"termin8", token.IDENT, "termin8"},
		{"_a2D", token.IDENT, "_a2D"},
		{"2a", token.ILLEGAL, "2a"},
		{"10", token.INT, "10"},
		{"=", token.ASSIGN, "="},
		{"+", token.PLUS, "+"},
		{"-", token.MINUS, "-"},
		{"!", token.BANG, "!"},
		{"*", token.ASTERISK, "*"},
		{"/", token.SLASH, "/"},
		{"<", token.LT, "<"},
		{">", token.GT, ">"},
		{"==", token.EQ, "=="},
		{"!=", token.NOT_EQ, "!="},
		{",", token.COMMA, ","},
		{";", token.SEMICOLON, ";"},
		{"(", token.LPAREN, "("},
		{")", token.RPAREN, ")"},
		{"{", token.LBRACE, "{"},
		{"}", token.RBRACE, "}"},
		{"fn", token.FUNCTION, "fn"},
		{"let", token.LET, "let"},
		{"true", token.TRUE, "true"},
		{"false", token.FALSE, "false"},
		{"if", token.IF, "if"},
		{"else", token.ELSE, "else"},
		{"return", token.RETURN, "return"},
	}

	allTokens := make(map[token.TokenType]bool)
	for _, v := range token.TokenTypeValues() {
		allTokens[v] = true
	}

	for _, tc := range cases {
		delete(allTokens, tc.expectedType)
	}
	if len(allTokens) > 0 {
		t.Fatalf("not all token types are tested: %v", allTokens)
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%s token", tc.expectedType.String()), func(t *testing.T) {
			if !tc.expectedType.IsATokenType() {
				t.Errorf("invalid token type: %d", tc.expectedType)
			}

			l := New(tc.token)
			tok := l.NextToken()

			if tok.Type() != tc.expectedType {
				t.Errorf("cases[%d]: expected token type %s, got %s",
					i, tc.expectedType.String(), tok.Type().String())
			}

			if tok.Literal() != tc.expectedLiteral {
				t.Errorf("cases[%d]: expected literal '%s', got '%s'",
					i, tc.expectedLiteral, tok.Literal())
			}
		})
	}
}

func TestNextToken_ActualScript(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
!-/*5;
5 < 10 > 5;

if (5 < 10) {
	return true;
} else {
	return false;
}

10 == 10;
10 != 9;
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type() != tt.expectedType {
			t.Fatalf("cases[%d]: expected token type %s, got %s",
				i, tt.expectedType.String(), tok.Type().String())
		}

		if tok.Literal() != tt.expectedLiteral {
			t.Fatalf("cases[%d]: expected literal '%s', got '%s'",
				i, tt.expectedLiteral, tok.Literal())
		}
	}
}
