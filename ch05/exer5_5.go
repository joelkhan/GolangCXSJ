/*
 * 若一个数恰好等于它的平方数的右端，则为同构数，
 * 如5的平方是25，5就是同构数，找出1~1000之间的同构数
 */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Printf("%t\n", isSameStructure(376))
	for i := 1; i <= 1000; i++ {
		if isSameStructure(i) {
			fmt.Printf("%10d %10d\n", i, i*i)
		}
	}
}

func isSameStructure(n int) (res bool) {
	a := n * n
	nStr := strconv.Itoa(n)
	aStr := strconv.Itoa(a)
	// fmt.Println(nStr)
	// fmt.Println(aStr)
	var i, j int
	for i, j = len(nStr)-1, len(aStr)-1; i >= 0; i, j = i-1, j-1 {
		if nStr[i] != aStr[j] {
			break
		}
	}
	if i == -1 {
		res = true
	} else {
		res = false
	}
	return
}
