package codeRule

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime/debug"
	"testing"
	"time"
)



func Foo1(a int ,b int)(string, bool){
	return "hello",true
}


func handler( w http.ResponseWriter, r *http.Request){
	io.WriteString(w,r.URL.Query().Get("param1"))
}

func Test04(t *testing.T){
	//http.HandlerFunc("/",handler)
	http.ListenAndServe(":8080",nil)
}

func Foo2(a int ,b int)(str string, ok bool){
	str = "hello"
	ok = true
	return
}



func Test02(t *testing.T) {
	defer func() {
		// fatal err 并不能被recover
		err := recover()
		fmt.Println(err)
	}()



	c := make(map[string]int)
	go func() {
		for j := 0; j < 1000000; j++ {
			c[fmt.Sprintf("%d ", j)] = j
		}
	}()

	go func() {
		for j := 0; j < 1000000; j++ {
			fmt.Println(c[fmt.Sprintf("%d", j)])
		}
	}()

	time.Sleep(time.Second * 60)
}

func Test01(t *testing.T) {
	go RecoverWrap(fool)()

	go foo2()
}

func foo2() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("SYSTEM ACTION PANIC: %v, stack: %v", string(debug.Stack()))
		}
	}()
}

func fool() {
	fmt.Println("fool")
}

func RecoverWrap(handle func()) func() {
	return func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("SYSTEM ACCTION PANIC : %v,stack %v", r, string(debug.Stack()))
			}
		}()
	}
}
