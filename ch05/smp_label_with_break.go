package main

import (
	"fmt"
)

func main() {
	label_with_break()
}

func label_with_break() {
LABEL1:
	for {
		for i := 0; i < 5; i++ {
			if i > 3 {
				break LABEL1
			}
			fmt.Println(i)
		}
	}
	fmt.Println("finished BREAK~")
}
