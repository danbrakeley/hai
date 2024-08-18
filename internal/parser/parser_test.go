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
				// "expected next token to be int, got semicolon instead",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			l := lexer.New(tc.input)
			p := New(l)

			program := p.ParseProgram()

			errors := p.Errors()
			if !checkErrors(t, errors, tc.errors) {
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

			for i, ident := range tc.expectedIdents {
				stmt := program.Statements[i]
				if stmt.TokenLiteral() != "let" {
					t.Fatalf("s.TokenLiteral not 'let', got %q", stmt.TokenLiteral())
				}

				letStmt, ok := stmt.(*ast.LetStatement)
				if !ok {
					t.Fatalf("expected *ast.Statement, got %T", stmt)
				}

				if letStmt.Name.Value != ident {
					t.Fatalf("letStmt.Name.Value not '%s', got %s", ident, letStmt.Name.Value)
				}

				if letStmt.Name.TokenLiteral() != ident {
					t.Fatalf("letStmt.Name.TokenLiteral() not '%s', got %s", ident, letStmt.Name.TokenLiteral())
				}
			}
		})
	}
}

func TestReturnStatements(t *testing.T) {
	cases := []struct {
		name          string
		input         string
		expectedCount int
		errors        []string
	}{
		{
			"single valid return",
			"return 5;",
			1,
			[]string{},
		},
		{
			"missing semicolon",
			"return 5",
			0,
			[]string{"expected next token to be semicolon, got eof instead"},
		},
		// {
		// 	"missing expression",
		// 	"return ;",
		// 	0,
		// 	[]string{"expected next token to be int, got semicolon instead",},
		// },
		{
			"three valid returns",
			`
return 5;
return 10;
return 993322;`,
			3,
			[]string{},
		},
		{
			"three invalid returns",
			`
return ;
return 10
return return;`,
			// 0,
			2,
			[]string{
				// "expected next token to be int, got semicolon instead",
				// "expected next token to be semicolon, got return instead",
				// "expected next token to be int, got return instead",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			l := lexer.New(tc.input)
			p := New(l)

			program := p.ParseProgram()

			errors := p.Errors()
			if !checkErrors(t, errors, tc.errors) {
				return
			}

			if program == nil {
				t.Fatalf("ParseProgram() returned nil")
			}
			if len(program.Statements) != tc.expectedCount {
				t.Fatalf("expected len(program.Statements) to be %d, got %d",
					tc.expectedCount, len(program.Statements),
				)
			}
		})
	}
}

func checkErrors(t *testing.T, actual, expected []string) bool {
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
