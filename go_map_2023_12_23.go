// go语言中map的定义和使用
// 跟python里的dictionary差不多
package main

import "fmt"

func main() {
	score := make(map[string]int, 10) //只声明的map变量是没有分配空间的，必须要通过make函数来分配空间
	score["张三"] = 0
	fmt.Println(score["张三"])

	id := map[string]int{ //也可以用这种方式来定义map
		"张三": 21,
		"李四": 20, //注意这种方式括号里的每行最后都要加上逗号
	}
	fmt.Println(id["张三"], "\t", id["李四"])
}
