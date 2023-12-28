// go语言中协程的创建和使用
package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("test()" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	go test() //创建协程的语法
	for i := 1; i <= 10; i++ {
		fmt.Println("main()" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
