package main

import "fmt"

func main()  {
	fmt.Println("hello,golang")

	var i="aaa"//go语言中定义变量以后必须要使用
    fmt.Println(i)
	fmt.Printf("%v\n",i)

	var a int=10
	var b int=3
	var c int=5
    fmt.Println("a=",a,"b=",b,"c=",c)
	fmt.Printf("a=%v b=%v c=%v",a,b,c)

}