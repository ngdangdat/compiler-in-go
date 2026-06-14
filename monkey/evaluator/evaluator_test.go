package evaluator

import (
	"testing"

	"github.com/ngdangdat/compiler-in-go/monkey/lexer"
	"github.com/ngdangdat/compiler-in-go/monkey/object"
	"github.com/ngdangdat/compiler-in-go/monkey/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5;", 5},
		{"10;", 10},
	}
	for _, tt := range tests {
		obj := testEval(tt.input)
		testIntegerObject(t, obj, tt.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	res, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer, got=%T (%v)", obj, obj)
		return false
	}
	if res.Value != expected {
		t.Errorf("object has wrong value, expected=%d got=%d", expected, res.Value)
		return false
	}
	return true
}
