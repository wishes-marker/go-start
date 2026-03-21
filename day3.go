//切片，切片里面没有长度
package main

import("fmt"
       "sort"
	)

func main(){
var nameList[]string 
//   nameList =append(nameList,"枫枫")
//   nameList =append(nameList,"张三")
//   fmt.Println(nameList[0]) 输出数组0【枫枫】
  fmt.Println(nameList == nil)

  ageList :=make([]int,3)
  fmt.Println(ageList)//int默认值0，输出【0 0 0】

  array:=[3]int{1,2,3}
  slices :=array[:]//切完整数组
  fmt.Println(slices)
  fmt.Println(array[0:2])//要前不要尾，输出【1 2】
  fmt.Println(array[1:2])//要前不要尾，输出【2】

//切片排序
  var ints = []int{4,8,3,2}
  sort.Ints(ints)//升序 
  fmt.Println(ints)
  sort.Sort(sort.Reverse(sort.IntSlice(ints)))//降序
  fmt.Println(ints)

}