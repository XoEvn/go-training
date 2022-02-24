/**
 * @Author: evnxo
 * @Description:
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2021/12/11 2:52 下午
 */

package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func echo(conn net.Conn) {
	defer conn.Close()
	b := make([]byte, 512)
	for {
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("client disconnected")
			break
		}
		if err != nil {
			log.Println("Unexpected error")
			break
		}
		log.Printf("Received %d bytes:%s\n", size, string(b))
		log.Println("Writing data")
		if _, err := conn.Write(b[:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listen on 0.0.0.0:20080")
	for {
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go echo(conn)
	}
	//var (
	//	reader FooReader
	//	writer FooWriter
	//)
	//
	//if _, err := io.Copy(&writer, &reader); err != nil {
	//	log.Fatalln("Unable to read/write data")
	//}
	//input := make([]byte, 4096)
	//s, err := reader.Read(input)
	//if err != nil {
	//	log.Fatalln("Unable to read data")
	//}
	//fmt.Printf("Read %d bytes from stdin \n", s)
	//
	//s, err = writer.Write(input)
	//if err != nil {
	//	log.Fatalln("Unable to write data")
	//}
	//fmt.Printf("Wrote %d bytes to stdout\n", s)
	//
	//ports := make(chan int, 100)
	//results := make(chan int)
	//var openports []int
	//for i := 0; i < cap(ports); i++ {
	//	go worker(ports, results)
	//}
	//go func() {
	//	for i := 1; i <= 1024; i++ {
	//		ports <- i
	//	}
	//}()
	//for i := 0; i < 1024; i++ {
	//	port := <-results
	//	if port != 0 {
	//		openports = append(openports, port)
	//	}
	//}
	//close(ports)
	//close(results)
	//sort.Ints(openports)
	//for _, port := range openports {
	//	fmt.Printf("%d open \n", port)
	//}
	//ports := make(chan int, 100)
	//var wg sync.WaitGroup
	//for i := 0; i < cap(ports); i++ {
	//	go worker(ports, &wg)
	//}
	//for i := 1; i <= 1024; i++ {
	//	wg.Add(1)
	//	ports <- i
	//}
	//wg.Wait()
	//close(ports)

	//----- 带顺序的执行net.dial
	//var wg sync.WaitGroup
	//for i := 1; i <= 1024; i++ {
	//	wg.Add(1)
	//	go func(j int) {
	//		defer wg.Done()
	//		address := fmt.Sprintf("scanme.nmap.org:%d", j)
	//		conn,err:=net.Dial("tcp",address)
	//		if err!=nil{
	//			return
	//		}
	//		conn.Close()
	//		fmt.Printf("%d open\n\n", j)
	//	}(i)
	//}
	//wg.Wait()
	//------
	//---基础的网络李连接
	//_, err := net.Dial("tcp", "scanme.nmap.org:80")
	//if err == nil {
	//	fmt.Println("Connection successful")
	//}
}

//func worker(ports chan int, wg *sync.WaitGroup) {
//	for p := range ports {
//		fmt.Println(p)
//		wg.Done()
//	}
//}
