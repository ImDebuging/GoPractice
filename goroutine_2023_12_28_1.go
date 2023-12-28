// go语言中协程和主线程的关系，“主死从随”
package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println(1)
	}() //可以使用匿名函数创建协程
	time.Sleep(time.Second) //创建协程以后，主线程要等待协程执行完毕，所以要等待一秒
}
