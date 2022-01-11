package main

import (
	"fmt"
	"strings"
)

func main() {
	smp3_14()
}

func smp3_14() {
	var str1, str2 string = "Go", "go"
	fmt.Printf("%t\n", strings.EqualFold(str1, str2))
}
