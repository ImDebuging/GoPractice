// go语言中面向对象编程思想的引入
package main

import "fmt"

type Teacher struct {
	Name   string
	Age    int
	School string
}

func (t Teacher) PrintAge() { //go语言中给类定义方法的格式，函数名在形参表之后，以此跟函数做区分
	fmt.Println(t.Age)
}

func (t *Teacher) ChangeSchool() { //go语言中如果想在方法中对结构体的某个属性做出改变，则需要通过指针传递的方式
	fmt.Scanln(&t.School)
}

func main() {
	zhong := Teacher{Name: "zhong", Age: 32, School: "Tsinghua"}
	fmt.Println(zhong.Name)
	zhong.PrintAge()
	zhong.ChangeSchool() //不过编译器会做出优化，传递的过程中可以省略地址操作符“&”
	fmt.Println(zhong)

	lu := Teacher{"luyan", 25, "JLZX"}
	fmt.Println(lu)
}
