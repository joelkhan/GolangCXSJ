package main

import (
	"fmt"
)

func main() {
	var cp1, cp2 complex64
	cp1 = 1.2 + 3.4i
	cp2 = cp1
	cp3 := complex(1.2, 3.4)
	fmt.Printf("%T, %v\n", cp1, cp1)
	fmt.Printf("%T, %v\n", cp2, cp2)
	fmt.Printf("%T, %v\n", cp3, cp3)
}
