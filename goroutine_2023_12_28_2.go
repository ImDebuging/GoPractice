package main

import (
	"fmt"
	"strconv"
	"sync"
)

var wg sync.WaitGroup // 定义了一个等待组

func main() {
	for i := 1; i <= 5; i++ {
		wg.Add(1) //将要创建新的协程时，等待组中协程数量加一
		go func(n int) {
			fmt.Println("协程" + strconv.Itoa(n))
			wg.Done() //一个协程将要结束时，等待组中协程数量减一
		}(i)
	}
	wg.Wait() //主线程等待所有协程结束后，利用这个函数自动结束
}
