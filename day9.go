//函数（指针）
//简单函数
package main

import "fmt"

// 基础无参函数
func sayHello() {
	fmt.Println("你好")
}

// 单参数函数
func param1(id string) {
	fmt.Println(id)
}

// 多参数同类型函数（注意：参数间用逗号分隔，类型只写一次）
func param2(id, userName string) {
	fmt.Println(id, userName)
}

// 可变参数求和函数（修正了原代码的非法变量名）
func add(numberList ...int) {
	var sum int
	for _, i2:= range numberList {
		sum += i2
	}
	fmt.Println(sum)
}
// 唯一的程序入口main函数
func main() {
	sayHello()
	param1("123456")
	param2("001", "张三")
	add(1, 2)
	add(1, 2, 3)
	add(1, 2, 3, 4)
}