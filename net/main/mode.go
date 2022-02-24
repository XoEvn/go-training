/**
 * @Author: evnxo
 * @Description:
 * @File:  mode.go
 * @Version: 1.0.0
 * @Date: 2021/12/11 4:35 下午
 */

package main

import (
	"fmt"
	"os"
)

type FooReader struct{}

func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Println("in > ")
	return os.Stdin.Read(b)
}

type FooWriter struct{}

func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Println("out > ")
	return os.Stdout.Write(b)
}
