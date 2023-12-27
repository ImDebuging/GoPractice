// 这个程序主要学了一下怎么在go中定义const常量
package main

import "fmt"

func main() {
	const LENGTH, WIDTH = 10, 5 //go中支持类似python的这种同时赋值方法
	var area int = LENGTH * WIDTH
	fmt.Printf("长为%d，宽为%d,面积为%d", LENGTH, WIDTH, area)
}
