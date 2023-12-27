// go语言中结构体的定义和初始化
package main

import "fmt"

type Book struct {
	price  int
	author string //结构体中的属性前面不用加关键字“var”，因为他们不是变量
}

func main() {
	var game_of_thrones Book
	game_of_thrones.price = 70
	game_of_thrones.author = "George R.R Martin"
	//fmt.Println(game_of_thrones)
	hongloumeng := Book{60, "曹雪芹"} //go语言中也可以采取这样的方式定义结构体变量，体会“:=”赋值运算符的好处
	fmt.Println(hongloumeng)
}
