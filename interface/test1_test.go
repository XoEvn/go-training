package _interface

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var printer Printer
		printer = FuncCaller(func(p interface{}) { fmt.Println(p) })
		printer.Print("Golang is good")
	})
}
