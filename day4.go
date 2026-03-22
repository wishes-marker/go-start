//map,创建时一定要赋值，一定要初始化
package main

	import"fmt"

	func main(){
    var userMap map[int]string=map[int]string{
		1:"枫枫",
		2:"张三",
		4:"",
	}
	fmt.Println(userMap)
	fmt.Println(userMap[1])
	fmt.Printf("%#v\n",userMap[3])
	fmt.Printf("%#v\n",userMap[4])
	value :=userMap[4]
	value,ok :=userMap[4]//判断是有的空字符串，还是无的空字符串
	fmt.Println(value,ok)

	userMap[1]="枫枫知道"
	fmt.Println(userMap)

	delete(userMap,4)//删除
	fmt.Println(userMap)

	var aMap=make(map[string]string)

	aMap["name"]="枫枫"
	fmt.Println(aMap)
 


}