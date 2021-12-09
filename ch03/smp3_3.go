package main

import (
	"fmt"
	"strings"
)

func main() {
	smp3_3()
}

func smp3_3() {
	id := 100
	name := "李明"
	grade := 91.5
	fmt.Printf("%v %v %v\n", id, name, grade)
	fmt.Printf("%#v %#v %#v\n", id, name, grade)
	fmt.Printf("%T %T %T\n", id, name, grade)
	yes, no := true, false
	fmt.Printf("%t %t\n", yes, no)
	fmt.Printf("%t\n", strings.EqualFold("Go", "go"))
}
