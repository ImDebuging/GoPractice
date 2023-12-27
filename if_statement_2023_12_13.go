//这个程序是我用来练习条件分支（if语句）的
/*
题目如下：
写一个程序，要求用户输入一个数字，然后判断该数字属于以下哪个范围：
如果数字小于0，则输出"输入的数字小于0"
如果数字等于0，则输出"输入的数字等于0"
如果数字大于0且小于10，则输出"输入的数字大于0且小于10"
如果数字大于等于10，则输出"输入的数字大于等于10"
请你完成这个程序。
*/
package main

import "fmt"

func main() {
	var num int
	fmt.Println("请输入一个数字: ")
	fmt.Scanln(&num)
	if num < 0 {
		fmt.Println("输入的数字小于0 ")
	} else if num == 0 {
		fmt.Println("输入的数字等于0 ")
	} else if num > 0 && num < 10 {
		fmt.Println("输入的数字大于0且小于10 ")
	} else {
		fmt.Println("输入的数字大于10 ")
	}
}
