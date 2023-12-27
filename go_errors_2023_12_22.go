// go语言中的错误处理机制，利用内置库errors的New函数自定义错误信息
package main

import (
	"errors"
	"fmt"
)

func main() {
	a, err := test(4, 0)
	if err == nil { //处理错误的模板
		fmt.Println(a)
	} else {
		fmt.Println(err)
	}
	fmt.Println("!!!")
}

func test(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Divided by zero") //利用内置库errors的New函数自定义错误信息
	} else {
		return a / b, nil
	}
}
