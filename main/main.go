package main

import (
	"fmt"
	"sync"
	"time"
)

func printer(b bool) {
	switch b {
	case true:
		fmt.Println("*")
	default:
		fmt.Println(" ")
	}
}

var m *sync.RWMutex

func init() {
	m = new(sync.RWMutex)
}

func read() {
	m.RLock()
	fmt.Println("startR")
	time.Sleep(time.Second)
	fmt.Println("endR")
	m.RUnlock()
}

func write() {
	m.Lock()
	fmt.Println("startW")
	time.Sleep(time.Second)
	fmt.Println("endW")
	m.Unlock()
}

func main() {

	go read()
	go read()
	go write()

	time.Sleep(time.Second * 3)

	//var mutex sync.Mutex
	//
	//wait := sync.WaitGroup{}
	//
	//now := time.Now()
	//
	//for i := 1; i <= 3; i++ {
	//
	//	wait.Add(1)
	//	go func(i int) {
	//		mutex.Lock()
	//		time.Sleep(time.Second)
	//		mutex.Unlock()
	//		defer wait.Done()
	//	}(i)
	//
	//}
	//wait.Wait()
	//
	//duration := time.Since(now)
	//fmt.Print(duration)
	//a := make(chan int)
	//close(a)
	////close(a)close
	//a = nil
	//close(a)

	//cs:=make(chan *os.File,100)

}
