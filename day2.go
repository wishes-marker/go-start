//数组，长度被限制死了，所以不经常使用
package main

import"fmt"

func main (){
	var nameList[3]string=[3]string{"枫枫","张三","李四"}
	fmt.Println(nameList)
	//"枫枫","张三","李四"
	//   0     1     2
	//  -3    -2    -1
	fmt.Println(nameList[0])
	fmt.Println(len(nameList))
	fmt.Println(nameList[len(nameList)-1])
	// fmt.Println(nameList[-1]) go语言不支持
}
