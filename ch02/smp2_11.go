package main

import (
	"fmt"
)

func main() {
	var i int
	var iPtr *int
	i = 100
	iPtr = &i
	fmt.Println(*iPtr)

	var c, d float32 = 2.5, 6.3
	fmt.Println(d - c)

	var t1 bool = true
	fmt.Println(!t1)

	fmt.Println(string("Golang"[0]))
}
