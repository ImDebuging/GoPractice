// go语言中实现复制一个文件的内容到其他文件的方法
package main

import (
	"fmt"
	"os"
)

func main() {
	file1Path := "test.txt"
	file2Path := "test2.txt"
	//实现的思路就是先读取要复制的文件的内容，然后写入目标文件

	file1, err1 := os.ReadFile(file1Path)
	if err1 != nil {
		panic(err1)
	} //读取要复制的文件内容

	err2 := os.WriteFile(file2Path, file1, 0666)
	if err2 != nil {
		panic(err2)
	} //写入目标文件

	file2, err3 := os.ReadFile(file2Path)
	if err3 != nil {
		panic(err3)
	}
	for _, v := range file2 {
		fmt.Printf(string(v))
	} //打印目标文件的内容来测试是否复制成功
}
