package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num = rand.Intn(10) + 1
	if num < 5 {
		fmt.Println(num)
	}
	num = rand.Intn(10) + 1
	fmt.Println(num)
}
