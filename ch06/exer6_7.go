package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var num, numEr int
	suffixEr := []byte("er")
	suffixEr2 := []byte("ER")
	buf := make([]byte, 1024)
	f, _ := os.Open("Lesson49.txt")
	defer f.Close()
	n, _ := f.Read(buf)
	fmt.Println(string(buf[:n]), n)
	for _, sentence := range bytes.FieldsFunc(buf[:n], f1) {
		for _, words := range bytes.Fields(sentence) {
			num++
			//fmt.Printf("%q", words)
			if bytes.HasSuffix(words, suffixEr) || bytes.HasSuffix(words, suffixEr2) {
				numEr++
				fmt.Printf("%q", words)
			}
		}
	}
	fmt.Println()
	fmt.Println("文件", f.Name(), "，共有", num, "个单词")
	fmt.Println("其中，以er结尾的单词共有", numEr, "个")
}

func f1(a rune) bool {
	if a == ':' || a == ',' || a == '.' || a == '!' {
		return true
	} else {
		return false
	}
}
