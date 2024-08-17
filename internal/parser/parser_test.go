package parser

import (
	"testing"

	"github.com/danbrakeley/hai/internal/ast"
	"github.com/danbrakeley/hai/internal/lexer"
)

func TestLetStatements(t *testing.T) {
	cases := []struct {
		name           string
		input          string
		expectedIdents []string
		errors         []string
	}{
		{
			"single valid let",
			"let x = 5;",
			[]string{"x"},
			[]string{},
		},
		{
			"missing semicolon",
			"let x = 5",
			[]string{},
			[]string{"expected next token to be semicolon, got eof instead"},
		},
		{
			"missing identifier",
			"let = 5;",
			[]string{},
			[]string{"expected next token to be ident, got assign instead"},
		},
		{
			"three valid lets",
			`
let x = 5;
let y = 10;
let foobar = 838383;`,
			[]string{"x", "y", "foobar"},
			[]string{},
		},
		{
			"three invalid lets",
			`
let = 5;
let y 10;
let foobar = ;`,
			[]string{},
			[]string{
				"expected next token to be ident, got assign instead",
				"expected next token to be assign, got int instead",
				"expected next token to be int, got semicolon instead",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			l := lexer.New(tc.input)
			p := New(l)

			program := p.ParseProgram()

			errors := p.Errors()
			if !assertErrors(t, errors, tc.errors) {
				return
			}

			if program == nil {
				t.Fatalf("ParseProgram() returned nil")
			}
			if len(program.Statements) != len(tc.expectedIdents) {
				t.Fatalf("expected len(program.Statements) to be %d, got %d",
					len(tc.expectedIdents), len(program.Statements),
				)
			}

			for i, tt := range tc.expectedIdents {
				stmt := program.Statements[i]
				if !assertLetStatment(t, stmt, tt) {
					return
				}
			}
		})
	}
}

func assertErrors(t *testing.T, actual, expected []string) bool {
	t.Helper()
	if len(actual) != len(expected) {
		var msg string
		if len(expected) == 0 {
			msg = "expected 0 parse errors, "
		} else {
			msg = "expected parse errors:\n"
			for _, e := range expected {
				msg += "\t" + e + "\n"
			}
		}
		if len(actual) == 0 {
			msg += "got none"
		} else {
			msg += "got:\n"
			for _, e := range actual {
				msg += "\t" + e + "\n"
			}
		}
		t.Error(msg)
	} else {
		for i := range actual {
			if actual[i] != expected[i] {
				t.Errorf("expected error %d to be:\n\t%s\ngot:\n\t%s", i, expected[i], actual[i])
			}
		}
	}
	return len(actual) == 0
}

func assertLetStatment(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let', got %q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("expected *ast.Statement, got %T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s', got %s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s', got %s", name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}
