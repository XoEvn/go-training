package exprDirectory

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"testing"
)


func Test001(t *testing.T){
	expr,_ := parser.ParseExpr(`1+2*3+x`)
	fmt.Println(Eval0(expr,map[string]float64{
		"x":100,
	}))
}

func Eval0(exp ast.Expr,vars map[string]float64) float64{
	switch exp :=exp.(type) {
	case *ast.BinaryExpr:
		return EvalBinaryExpr0(exp,vars)
	case *ast.BasicLit:
		f,_ :=strconv.ParseFloat(exp.Value, 64)
		return f
	case *ast.Ident:
		return vars[exp.Name]
	}
	return 0
}



func EvalBinaryExpr0(exp *ast.BinaryExpr,vars map[string]float64)float64{
	switch exp.Op {
	case token.ADD:
		return Eval0(exp.X,vars) + Eval0(exp.Y, vars)
	case token.MUL:
		return Eval0(exp.X,vars) * Eval0(exp.Y,vars)
	}
	return 0
}
