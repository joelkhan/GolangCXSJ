package main

import (
	"fmt"
)

func main() {
	smp4_6()
	fmt.Println(exampleX(0))
	fmt.Println(exampleX(3))
}

func smp4_6() {
	a := 10
	if a := 1; a > 0 {
		fmt.Println(a)
	}
	fmt.Println(a)
}

func exampleX(x int) int {
	if x == 0 {
		return 10
	} else {
		return x
	}
}
