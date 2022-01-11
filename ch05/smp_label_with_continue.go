package main

import (
	"fmt"
)

func main() {
	label_with_continue()
}

func label_with_continue() {
LABEL1:
	for i := 0; i < 5; i++ {
		fmt.Println("UP:", i)
		for {
			continue LABEL1
			fmt.Println("MIDDLE of Never.")
		}
		fmt.Println("DOWN:", i)
	}
	fmt.Println("finished CONTINUE~")
}
