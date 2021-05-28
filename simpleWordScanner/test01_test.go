package simpleWordScanner

import (
	"fmt"
	"go/scanner"
	"go/token"
	"testing"
)

func Test02(t *testing.T){
	a := token.Position{Filename: "hello.go",Line:1,Column: 2}
	ab := token.Position{Filename: "hello.go",Offset:2,Line:1,Column: 2}
	b := token.Position{Filename: "hello.go",Line:1}
	c := token.Position{Filename: "hello.go"}
	d := token.Position{Line: 1,Column: 2}
	e := token.Position{Line: 1}
	f := token.Position{Column: 2}

	fmt.Println(a.String())
	fmt.Println(b.String())
	fmt.Println(c.String())
	fmt.Println(d.String())
	fmt.Println(e.String())
	fmt.Println(f.String())
	fmt.Println(ab.String())
}


func Test01(t *testing.T){
	var src = []byte (`println("你好,世界")"`)
	var fset = token.NewFileSet()

	var file = fset.AddFile("hello.go",fset.Base(),len(src))

	var s scanner.Scanner

	s.Init(file,src,nil,scanner.ScanComments)

	for {
		pos, tok, lit :=s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n",fset.Position(pos),tok,lit)
	}
}