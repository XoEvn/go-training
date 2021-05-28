package faceValue

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func Test02(t *testing.T){
	expr,_ :=parser.ParseExpr(`"9527"`)
	ast.Print(nil,expr)
}


func Test01(t *testing.T){
	var lit9527 = &ast.BasicLit{
		Kind: token.INT,
		Value: "8527",
	}
	ast.Print(nil,lit9527)
}