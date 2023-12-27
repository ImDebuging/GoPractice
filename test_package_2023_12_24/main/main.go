// go语言中结构体定义及跨包（不同的package）里使用
package main

import (
	"fmt"
	m "go_-practice/test_package_2023_12_24/model" //Student这个结构体是在model文件夹下的model.go这个文件里定义的
)

func main() {
	stu1 := m.Student{12, "张三"}
	fmt.Println(stu1)
}
