//for循环
package main

import "fmt"

func main(){
	//传统for循环
	// var sum int
	// for i:=1;i<=100;i++{//for 初始化；条件；操作{}
	// 	sum+=i
	// }
	// 	fmt.Println(sum)

	//while模式
	// var sum int 
	// var i int=1
	// for i <=100{//判断条件
	// 	sum+=i//循环体里进行条件变化
	// 	i++
	// }
	// fmt.Println(sum)

	//do while模式，先执行一次循环体，再判断
	// i:=0
	// sum :=0
	// for {
	// 	sum +=i
	// 	i++
	// 	if i==101{
	// 		break
	// 	}
	// }
	// 	fmt.Println(sum)

	//遍历，循环切片
	// var list =[]string{"枫枫","知道"}
	// for i :=0;i<len(list);i++{
	// 	fmt.Println(i,list[i])
	// }
	// for index,item:=range list{
	// 	fmt.Println(index,item)
	// }

	//循环map
	var userMap=map[string]string{"name":"枫枫"}
	for key,value :=range userMap{
		fmt.Println(key,value)
	}

}