package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("sample code 5-2:")
	smp5_2()
	fmt.Println()
	fmt.Println("my try:")
	my_try()
}

func smp5_2() {
	var row, col int = 10, 5
	var seatNo string
	for k := 0; k < row; k++ {
		for i, j := k*col, (k+1)*col-1; i < (k+1)*col && j >= k*col; i, j = i+1, j-1 {
			if k%2 == 0 {
				seatNo += " " + strconv.Itoa(i)
			} else {
				seatNo += " " + strconv.Itoa(j)
			}
		}
		fmt.Printf("%s\n", seatNo)
		seatNo = ""
	}
}

func my_try() {
	var row, col = 10, 5
	var num = 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i%2 == 1 {
				numMax := (i+1)*col - 1
				numMin := i * col
				numTotal := numMax + numMin
				fmt.Printf(" %2d", numTotal-num)
			} else {
				fmt.Printf(" %2d", num)
			}
			num++
		}
		fmt.Println()
	}
}
