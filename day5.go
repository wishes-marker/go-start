//判断语句if语句
package main

import "fmt"

func main(){
	var age int
	fmt.Println("请输入你的年龄：")
	fmt.Scan(&age)//使控制台可读取年龄
	//下面是中断式 卫语句

	if age <=0{
		fmt.Println("未出生")
		return
	}
    if age <=18{
		fmt.Println("未成年")
		return
	}
	if age <=35{
		fmt.Println("青年")
		return
	}
	fmt.Println("中年")
	//下面是嵌套语句
	// if age <=18{
	// 	if age<=0{
	// 	   fmt.Println("未出生")
    //     }else{
	// 	   fmt.Println("未成年")
	// 	}
	// }else{
	// 	if age <=35{
	// 		fmt.Println("青年")
	// 	}else {
	// 		fmt.Println("中年")
	// 	}
	// }
	//下面是多条件式
	// if age <=0{
    //    fmt.Println("未出生")
	// }
	// if age <=18 && age >0{
    // fmt.Println("未成年")
	// }
    // if age <=35 && age >18{
    // fmt.Println("青年")
	// }
	// if age >35{
	// fmt.Println("中年")	
	// }

  
} 