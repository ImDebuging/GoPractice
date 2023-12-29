// go语言中管道的定义和使用
package main

import "fmt"

func main() {
	var ch chan int
	ch = make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("len:", len(ch))
}
