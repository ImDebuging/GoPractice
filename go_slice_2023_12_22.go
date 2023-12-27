// go语言中切片的定义和使用
package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	s := arr[1:3] //切片是基于数组的数据结构
	fmt.Println(s)
	fmt.Printf("%T", s)      //切片是[]int类型的数据
	s1 := make([]int, 4, 20) //go语言中也可以用make函数直接创建一个切片
	fmt.Println(s1)          //这个切片初始值为[]int[0 0 0 0]
	s1 = append(s1, 6, 7)    //利用append方法可以在切片后面添加元素，如何把添加后的切片返回
	fmt.Println(s1)          //如果想对原先的切片做出修改，则可用原先的切片去接受这个返回值

	s2 := []int{1, 2, 3, 4, 5}
	s3 := make([]int, 5, 10)
	copy(s3, s2) //利用copy函数可以将一个切片复制给另一个新的切片
	fmt.Println(s2, "\n", s3)

}
