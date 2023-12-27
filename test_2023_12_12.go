package main

import "fmt"

const (
	day      float64 = 28
	distance float64 = 56000000
)

var speed float64

func main() {
	speed = distance / day
	fmt.Println(speed)
}
