package main

import (
	"fmt"
)

func main() {
	a := []byte{'1', '2', '3', '4', 'z', 'x', '5', '6', 'd', 'g', '7', '8', 'D'}
	var index, x, start int
	fmt.Println("请输入起始索引位置……")
	fmt.Scanf("%d", &start)
	for {
		index, x = nextInt(a, start)
		//index, x = nextInt2(a, start)
		fmt.Printf("找到整数%d，下一个数的位置从%d开始。\n", x, index)
		if index != 0 && index < len(a) {
			start = index
			//continue
		} else {
			break
		}
	}
}

// 我的实现
func nextInt(b []byte, i int) (index, x int) {
	var findDigit bool
	for j := i; j < len(b); j++ {
		if b[j] >= '0' && b[j] <= '9' {
			if x <= 0 { // 第一次，找到第一个数字
				x = x*10 + int(b[j]-'0')
				findDigit = true
			} else if findDigit { // 第一次，找到连续数字
				x = x*10 + int(b[j]-'0')
			} else { // 第二次，找到第一个数字
				index = j
				break
			}
		} else {
			findDigit = false
		}
	}
	return
}

// 教材示例
func nextInt2(b []byte, i int) (index, x int) {
	var start bool
	if b[i] >= '0' && b[i] <= '9' {
		start = true
	} else {
		start = false
	}

	for ; i < len(b); i++ {
		if b[i] >= '0' && b[i] <= '9' {
			if start == false {
				index = i
				break
			} else {
				x = x*10 + int(b[i]) - '0'
			}
		} else if b[i] < '0' || b[i] > '9' {
			start = false
		}
	}
	return index, x
}
