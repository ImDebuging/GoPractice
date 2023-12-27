// 对go语言中的break关键字进行练习
package main

import "fmt"

func main() {
	//求1~100的和，当结果第一次大于300时，跳出循环并打印当前值
	sum, i := 0, 1
	for ; i <= 100; i++ {
		sum += i
		if sum > 300 {
			break
		}
	}
	fmt.Println("i=", i, ";", "sum=", sum)
}
