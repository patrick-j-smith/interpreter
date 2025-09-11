package parser

import (
	"interpreter/ast"
	"interpreter/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foobar = 838383;
	`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram returned nil program")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. "+
			"Instead got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		state := program.Statements[i]
		if !testLetStatement(t, state, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. Instead got=%q",
			s.TokenLiteral())
		return false
	}

	letState, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. Instead got=%T", s)
		return false
	}

	if letState.Name.Value != name {
		t.Errorf("letState.Name.Value not '%s'. Instead got=%s", name,
			letState.Name.Value)
		return false
	}

	if letState.Name.TokenLiteral() != name {
		t.Errorf("letState.Name.TokenLiteral() not '%s'. Instead got=%s", name,
			letState.Name.TokenLiteral())
		return false
	}

	return true
}
