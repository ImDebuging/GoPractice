//这个程序是我用来练习go语言中循环结构的
//go语言中只有for循环，没有类似C，C++中的while循环（感觉不太好）

package main

import "fmt"

func main() {
	var sum =0
	for var i = 1; i <= 5; i++ {
		sum += i
	}
	fmt.Println("1+2+3+4+5 =", sum)
}
