package main

import (
	"fmt"
)

const (
	a, b = 1, 2
	c, d
)

func main() {
	fmt.Println(a, b, c)
	fmt.Println()

	str := "Hello 世界!"
	n := len(str)
	fmt.Println("字节遍历：")
	for i := 0; i < n; i++ {
		ch := str[i]
		fmt.Printf("str[%d]=%v\n", i, ch)
	}
	fmt.Println()

	fmt.Println("字符遍历：")
	for i, ch := range str {
		fmt.Printf("str[%d]=%v\n", i, ch)
	}

	var bSlice = []byte{'H', 'e', 'l', 'l', 'o'}
	str1 := string(bSlice)
	fmt.Println(str1)

	/*
	 * Unicode是针对字符的编码方案，UTF-8是针对编码的存储方案。
	 * 字符的编码是“byte”或者“rune”来表示的，
	 * 字符（字符串）的存储是“string”来表示的。
	 */
	var rSlice = []rune{'H', 'e', 'l', 'l', 'o', '好'}
	str2 := string(rSlice)
	fmt.Println(str2)
	fmt.Println("字节遍历：")
	for i := 0; i < len(str2); i++ {
		ch := str2[i]
		fmt.Printf("str2[%d]=%v - %T\n", i, ch, ch)
	}
	fmt.Println()

	fmt.Println("字符遍历：")
	for i, ch := range str2 {
		fmt.Printf("str2[%d]=%v - %T\n", i, ch, ch)
	}
}
