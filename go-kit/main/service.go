package main


type IHelloWorld interface {
	say() string
}


type HelloWorld struct {}

func (h HelloWorld) say() string {
	return "hello world"
}