package main

import (
	"fmt"
)

func main() {
	funcEx2_5()
	funcEx2_8()
	funcEx2_9()
}

func funcEx2_5() {
	var a, b int = 15, 6
	var f float64
	f = float64(a / b)
	fmt.Println(f)
}

func funcEx2_8() {
	var a, b int = 10, 1
	var p *int
	p = &a
	a = *p + b
	fmt.Println(a)
}

func funcEx2_9() {
	var str1, str2 string
	fmt.Println("please input str1: ")
	fmt.Scanf("%s", &str1)
	fmt.Println("please input str2: ")
	fmt.Scanf("%s", &str2)
	fmt.Println("final string: ", str1+str2)
	var str3, str4 string
	fmt.Println("please input str3 str4: ")
	fmt.Scanf("%s %s\n", &str3, &str4)
	fmt.Printf("%#v && %#v\n", str3, str4)
}
