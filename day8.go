//break和continue
package main

import"fmt"

func main(){
	// for i:=1;i<=9;i++{
	// 	for j:=1;j<=i;j++{
	// 		fmt.Printf("%d*%d=%d\t",i,j,i*j)
	// 	}
	// 	fmt.Println()
	// }//打印出九九乘法表

	for i:=0;i<=10;i++{
		
		if i==5{
			// continue//跳过本次循环
			break//中止此处及以后循环
		}
	fmt.Println(i)
    }
}