package exprDirectory

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"testing"
)





func Eval(exp ast.Expr) float64 {
	switch exp := exp.(type) {
	case *ast.BinaryExpr:
		return EvalBinaryExpr(exp)
	case *ast.BasicLit:
		f,_ :=strconv.ParseFloat(exp.Value,64)
		return f
	}
	return 0
}

func EvalBinaryExpr(exp *ast.BinaryExpr) float64 {
	switch exp.Op {
	case token.ADD:
		return Eval(exp.X) + Eval(exp.Y)
	case token.MUL:
		return Eval(exp.X) * Eval(exp.Y)
	}
	return 0
}

func Test02(t *testing.T) {
	expr,_ := parser.ParseExpr(`1+2*3`)
	fmt.Println(Eval(expr))
}

func Test01(t *testing.T) {
	expr, _ := parser.ParseExpr(`1+2*3`)
	ast.Print(nil, expr)
}
