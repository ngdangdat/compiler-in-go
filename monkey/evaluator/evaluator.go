package evaluator

import (
	"github.com/ngdangdat/compiler-in-go/monkey/ast"
	"github.com/ngdangdat/compiler-in-go/monkey/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}
	return nil
}
