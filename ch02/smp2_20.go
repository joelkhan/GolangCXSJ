package main

import (
	"fmt"
)

type ByteSize float64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Printf("1KB = %.f Byte\n", KB)
	fmt.Printf("1MB = %.f Byte\n", MB)
	fmt.Printf("1GB = %.f Byte\n", GB)
	fmt.Printf("1TB = %.f Byte\n", TB)
	fmt.Printf("1PB = %.f Byte\n", PB)
	fmt.Printf("1EB = %.f Byte\n", EB)
	fmt.Printf("1ZB = %.f Byte\n", ZB)
	fmt.Printf("1YB = %.f Byte\n", YB)
}
