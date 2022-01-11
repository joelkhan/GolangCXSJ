package main

import (
	"fmt"
)

func main() {
	label_with_goto()
}

func label_with_goto() {
	for {
		for i := 0; i < 5; i++ {
			if i > 3 {
				goto LABEL1
			}
			fmt.Println(i)
		}
	}
LABEL1:
	fmt.Println("finished GOTO~")
}
