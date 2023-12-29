// go语言中利用管道和协程实现生产者消费者问题
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func writeData(ch chan int) {
	defer wg.Done()
	for i := 1; i <= 50; i++ {
		ch <- i
		fmt.Println("writeData", i)
		time.Sleep(time.Second)
	}
}

func readData(ch chan int) {
	defer wg.Done()
	for v := range ch { //go语言中遍历管道要用for-range循环来写，且管道没有索引
		fmt.Println("readData", v)
		time.Sleep(time.Second)
	}
}
func main() {
	var ch chan int
	ch = make(chan int, 50)
	wg.Add(2)
	go writeData(ch)
	go readData(ch)
	wg.Wait()
}
