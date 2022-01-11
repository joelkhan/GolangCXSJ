/*
 * 某铁路线上共有10个车站，如果每两个车站之间需要一种车票，问共需准备多少种车票
 */
package main

import (
	"fmt"
)

const platform = 10

func main() {
	var cntr int
	for i := 1; i <= platform-1; i++ {
		for j := i + 1; j <= platform; j++ {
			fmt.Printf("%2d --> %2d\n", i, j)
			cntr++
		}

	}
	fmt.Printf("total platform: %d\n", platform)
	fmt.Printf("total tickets: %d\n", cntr)
}
