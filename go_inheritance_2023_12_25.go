// go语言中继承的用法
package main

import (
	"fmt"
)

type Animal struct {
	Name string
	Age  int
}

func (a Animal) SayHello() {
	fmt.Println("Hello, I am", a.Name)
}

type Cat struct {
	Animal
}

func (c Cat) Meow() {
	fmt.Println("Meow")
}

func main() {
	c := Cat{}
	c.Animal.Name = "Tom"
	c.SayHello()
	c.Meow()
}
