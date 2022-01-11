package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Println("please input your string:")
	fmt.Scanf("%s\n", &str)
	fmt.Println(str)
	fmt.Println(len(str)) // 串的字节长度
	var rSlice = []rune(str)
	fmt.Println(len(rSlice))
}
