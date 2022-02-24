/**
 * @Author: evnxo
 * @Description:
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2021/12/28 8:50 下午
 */

package main

import (
	"fmt"
	"sync"
	"time"
)

var m *sync.RWMutex


func read(){
	m.RLock()

	fmt.Println("startR")
	
}

func main() {
	var mutex sync.Mutex

	wait := sync.WaitGroup{}

	now := time.Now()

	for i := 1; i <= 3; i++ {
		wait.Add(1)
		go func(i int) {
			mutex.Lock()
			time.Sleep(time.Second)
			mutex.Unlock()
			defer wait.Done()
		}(i)
	}
	wait.Wait()
	duration:=time.Since(now)
	fmt.Print(duration)
}

