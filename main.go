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

//类型推导方式定义变量
	d :=10
	e :=20
	f :=30
    fmt.Printf("d=%v e=%v f=%v \n",d,e,f)
	//使用Printf打印一个变量的类型
	fmt.Printf("d=%v d的类型是%T",d,d)
    

}