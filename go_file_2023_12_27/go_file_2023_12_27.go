// go语言中对文件的操作
package main

import (
	"fmt"
	"os" //用io流的方式对文件进行操作，搭建程序和数据源之间的io桥
)

func main() {
	//打开文件
	file, err := os.Open("test.txt")
	content, err := os.ReadFile(file.Name())
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	//没有出错，输出文件
	fmt.Println(string(content))
	defer file.Close()
}
