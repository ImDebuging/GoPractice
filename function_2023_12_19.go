// 从这个程序开始学习go语言中函数的定义
package main

import "fmt"

//go语言中函数定义的模板
/*
func 函数名(形参列表)(返回值列表){
	执行语句
	return + 返回值列表
}
*/

func Sum(a, b int) int { //注意，形参列表中不要加关键字var！！！
	return a + b
}

func SumAndSubstract(a int, b int) (int, int) { //go语言中支持多个返回值，返回值的类型都要写在返回值列表里
	return a + b, a - b
}

func Exchange(a *int, b *int) { //也可以没有返回值，类似C++中的void
	var c = *b
	*b = *a
	*a = c
}

func PrintArray(args ...int) { //go语言允许形参个数不确定
	for i := 0; i < len(args); i++ {
		fmt.Printf("%d", args[i])
	}
}

func main() {
	a, b := SumAndSubstract(20, 10)
	fmt.Printf("sum=%d,difference=%d\n", a, b)
	Exchange(&a, &b)
	fmt.Println(a, b)
	PrintArray(1, 2, 3, 4, 5, 6)
}
