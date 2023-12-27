// go语言中接口的定义
package main

import (
	"fmt"
)

type SayHello interface { //定义接口用interface关键字
	sayHello() //接口内部放置没有实现的方法
}

type Chinese struct {
}

func (c Chinese) sayHello() { //每个结构体中都要实现接口中的方法
	fmt.Println("你好")
}

type American struct {
}

func (a American) sayHello() { //每个结构体中都要实现接口中的方法
	fmt.Println("hello")
}

func greet(s SayHello) { //定义一个接口函数，接收一个接口类型参数
	s.sayHello()
}
func main() {
	greet(American{}) //根据传入结构体类型的不同，调用不同的接口函数
	greet(Chinese{})
}
