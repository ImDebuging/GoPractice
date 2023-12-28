// go语言，多个协程间的数据竞争(DataRace)
package main

import "sync"

var Num int = 0
var wg sync.WaitGroup

func Add() {
	defer wg.Done()
	for i := 0; i < 1000000; i++ {
		Num += 1
	}
}

func Sub() {
	defer wg.Done()
	for i := 0; i < 1000000; i++ {
		Num -= 1
	}
}

func main() {
	wg.Add(2)
	go Add()
	go Sub()
	wg.Wait()
	println(Num) //结果预期是0，但是由于Add协程和Sub协程之间存在数据竞争，导致结果不为0
}
