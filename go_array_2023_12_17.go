// go语言中数组的定义和初始化，顺便复习了go语言中for循环的语法
package main

import "fmt"

func main() {
	var arr [5]int
	for i := 0; i < 5; i++ {
		arr[i] = i
	}
	sum := 0
	arr2 := []int{1, 2, 3, 4, 5, 6}
	for _, v := range arr {
		sum += v
	}
	fmt.Println(sum, arr2)
}
