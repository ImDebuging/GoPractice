package main

import "fmt"

func main() {
	str := "Hello World!"
	// 遍历字符串中的每个字符
	for _, v := range str {
		// 打印每个字符
		fmt.Printf("%c", v)
	}
}
