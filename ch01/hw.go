package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello, world~")
	fmt.Println("你好，世界！")

	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.String())
	fmt.Println(t.Format("2006year 01month 02day"))
	fmt.Println(t.Weekday().String())

	/*
		var ch1, ch2 byte
		ch1 = 65
		ch2 = 'A'
		fmt.Println(string(ch1))
		fmt.Println(ch2)

		var rn1 rune
		rn1 = '中'
		rn2 := 22269    // 国
		rn3 := '\u56fd' // 国
		rn4 := 0x56fd   // 国
		fmt.Println(rn1)
		fmt.Println(string(rn1) + string(rn2))
		fmt.Println(string(rn3))
		fmt.Println(string(rn4))
	*/
}
