package main

import (
	"fmt"
)

func main() {
	var f1, f2 float32
	f1 = 12356.789e5
	f2 = f1 + 20
	fmt.Println(f2)

	var iPointer *int
	fmt.Println(iPointer)
	fmt.Printf("%v\n", iPointer)
	fmt.Printf("%#v\n", iPointer)
	fmt.Printf("%T\n", iPointer)
	fmt.Printf("%p\n", iPointer)
	fmt.Printf("%#p\n", iPointer)
}
