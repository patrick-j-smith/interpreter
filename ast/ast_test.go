package ast

import (
	"testing"

	"interpreter/token"
)

func TestString(t *testing.T) {

	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "Var1"},
					Value: "Var1",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "Var2"},
					Value: "Var2",
				},
			},
		},
	}

	if program.String() != "let Var1 = Var2;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}

}
