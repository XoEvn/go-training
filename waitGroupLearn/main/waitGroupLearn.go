/**
 * @Author: evnxo
 * @Description:
 * @File:  waitGroupLearn.go
 * @Version: 1.0.0
 * @Date: 2021/8/25 10:43 下午
 */

package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		time.Sleep(1*time.Second)
		fmt.Println("first G ")
		wg.Done()
	}()

	go func() {
		time.Sleep(1*time.Second)
		fmt.Println("second G")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("all done")
}
