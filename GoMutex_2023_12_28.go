// go语言中，利用互斥锁以避免数据竞争
package main

import "sync"

var Num int = 0
var wg sync.WaitGroup
var mtx sync.Mutex

func Add() {
	defer wg.Done()
	for i := 0; i < 1000000; i++ {
		mtx.Lock()
		Num += 1
		mtx.Unlock()
	}
}

func Sub() {
	defer wg.Done()
	for i := 0; i < 1000000; i++ {
		mtx.Lock()
		Num -= 1
		mtx.Unlock()
	}
}

func main() {
	wg.Add(2)
	go Add()
	go Sub()
	wg.Wait()
	println(Num) //加上了互斥锁mtx后，避免了数据竞争，Num的结果符合预期值，是0
}
