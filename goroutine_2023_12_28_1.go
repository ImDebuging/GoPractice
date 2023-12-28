// go语言中协程和主线程的关系，“主死从随”
package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	for i := 1; i <= 5; i++ {
		go func(n int) { //可以使用匿名函数创建协程}
			fmt.Println("协程" + strconv.Itoa(n))
		}(i) //为了达到预期效果（一个协程对应一个id），要避免这里的资源共享问题
	}
	time.Sleep(time.Second) //创建协程以后，主线程要等待协程执行完毕，所以要等待一秒
}
