// 这个程序是我用来练习go语言中多分支语句（switch语句）的
// 感觉go中的switch语句和C语言中的用法比较类似
package main

import "fmt"

//以下的程序实现了输入成绩marks，根据marks给出等级grade的功能
func main() {
	var (
		marks int
		grade string
	)
	fmt.Println("请输入您的成绩")
	fmt.Scanln(&marks)
	switch marks {
	case 90:
		grade = "A"
		fallthrough
	case 80:
		grade = "B"
		fallthrough
	case 70:
		grade = "C"
		fallthrough
	default:
		grade = "D"
	}
	fmt.Printf("您的分数是%d,等级是%v", marks, grade)
}
