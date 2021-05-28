package faceIdentifier

import (
	"go/ast"
	"go/parser"
	"testing"
)


func Test02(t *testing.T){
	expr,_ := parser.ParseExpr(`X`)
	ast.Print(nil,expr)
}

func Test01(t *testing.T){
	ast.Print(nil,ast.NewIdent(`X`))
}
