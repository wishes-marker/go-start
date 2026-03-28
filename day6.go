//判断语句switch语句，条件语句中go语言中会自动中止（break），想往下循环需要fallthrough
package main

import  "fmt"

func main(){
	var week int
	fmt.Printf("请输入星期（数字）：")
	fmt.Scan(&week)

	switch week {
	case 1:
		fmt.Println("周一")						
	case 2:
		fmt.Println("周二")
	case 3:
		fmt.Println("周三")
	case 4:
		fmt.Println("周四")
	case 5:
		fmt.Println("周五")	
	case 6,7:
		fmt.Println("周末")
	default:
		fmt.Println("错误的")
		
	}



}