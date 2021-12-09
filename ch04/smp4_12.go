package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf(t.Format("2006年 01月 02日\n"))
	switch t.Weekday().String() {
	case "Sunday":
		fmt.Printf("星期天\n")
	case "Monday":
		fmt.Printf("星期一\n")
	case "Tuesday":
		fmt.Printf("星期二\n")
	case "Wednesday":
		fmt.Printf("星期三\n")
	case "Thursday":
		fmt.Printf("星期四\n")
	case "Friday":
		fmt.Printf("星期五\n")
	case "Saturday":
		fmt.Printf("星期六\n")
	}

	ex4_5()
}

func ex4_5() {
	var n int
	fmt.Printf("请输入1～7的一个整数：")
	fmt.Scanf("%d", &n)
	//fmt.Printf("your number is: %d\n", n)

	switch n {
	case 7:
		fmt.Printf("星期天\n")
	case 1:
		fmt.Printf("星期一\n")
	case 2:
		fmt.Printf("星期二\n")
	case 3:
		fmt.Printf("星期三\n")
	case 4:
		fmt.Printf("星期四\n")
	case 5:
		fmt.Printf("星期五\n")
	case 6:
		fmt.Printf("星期六\n")
	default:
		fmt.Printf("请输入1～7的一个整数！\n")
	}

}
