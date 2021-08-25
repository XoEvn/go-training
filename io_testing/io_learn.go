/**
 * @Author: evnxo
 * @Description:
 * @File:  io_learn
 * @Version: 1.0.0
 * @Date: 2021/7/3 3:07 下午
 */

package io_testing

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func readBigFile(filePath string) error {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}

	buf := bufio.NewReader(f)
	count := 0
	for {
		count += 1
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			return err
		}
		fmt.Println("line", line)
		if count > 100 {
			break
		}
	}
	return nil
}

func readByFile() {
	data, err := ioutil.ReadFile("/Users/evnxo/12306/12306/tkcode.png")
	if err != nil {
		log.Fatal("err:", err)
		return
	}
	fmt.Println("data", string(data))
}

func writeFile() {
	err := ioutil.WriteFile("/Users/evnxo/1.txt/write_test.txt", []byte("Hello world!"), 0644)
	if err != nil {
		panic(err)
		return
	}
}
