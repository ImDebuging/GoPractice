// 今天学习了go语言中package的用法
// TMD,改bug累死老子了
//写package要注意这样几点“
/*
  1.用文件夹装你写的go package文件
  2.文件夹的名字尽可能与你定义的package文件名字保持一致
  3.package中定义的函数首字母要大写（为了让其他的包能引用这个package里的函数）
  4.调用的时候跟pyhton类似
*/
package main

import (
	"fmt"
	"go_-practice/my_package/test"
) //包含多个package的时候，最好使用这种格式

func main() {
	fmt.Println("111")
	test.Test()
}
