package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	copyLen, err := copyFile("dst.txt", "src.txt")
	if err != nil {
		return
	} else {
		fmt.Println(copyLen)
	}
}

func copyFile(dstName, srcName string) (copyLen int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
