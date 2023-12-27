//go语言中没有while循环，但是能用for循环达成类似的效果

package main

import "fmt"

func main() {
	flag := true
	for flag == true {
		fmt.Println("死循环")
	} //这样达成了while循环的效果
}
