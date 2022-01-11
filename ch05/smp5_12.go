package main

import (
	"fmt"
)

func main() {
	var str string
	var charCntr, numCntr, otherCntr int
	fmt.Println("请输入字符串……")
	fmt.Scanf("%s", &str)
	for _, v := range str {
		if v >= 48 && v <= 56 {
			numCntr++
		} else if (v >= 65 && v <= 90) || (v >= 97 && v <= 122) {
			charCntr++
		} else {
			otherCntr++
		}

	}
	fmt.Printf("数字: %d\n", numCntr)
	fmt.Printf("字母: %d\n", charCntr)
	fmt.Printf("其它: %d\n", otherCntr)
}
