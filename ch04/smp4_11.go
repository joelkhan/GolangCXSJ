package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, disc, x1, x2, p, q float64
	var cpx1, cpx2 complex128
	fmt.Println("请输入三个系数：")
	fmt.Scanf("%f, %f, %f\n", &a, &b, &c)
	if a == 0 {
		fmt.Printf("系数a为0，不是一元二次方程。\n")
	} else {
		disc = b*b - 4*a*c
		p = -b / (2 * a)
		if disc == 0 {
			x1 = p
			x2 = p
			fmt.Printf("判别式等于0，有两个相等的实根：x1 = x2 = %8.2f\n", p)
		} else if disc > 0 {
			q = math.Sqrt(disc) / (2 * a)
			x1 = p + q
			x2 = p - q
			fmt.Printf("判别式大于0，有两个不等的实根：x1 = %8.2f x2 = %8.2f\n", x1, x2)
		} else if disc < 0 {
			q = math.Sqrt(math.Abs(disc)) / (2 * a)
			cpx1 = complex(p, q)
			cpx2 = complex(p, -q)
			fmt.Printf("判别式小于0，有两个共轭的复根：cpx1 = %8.2f cpx2 = %8.2f\n", cpx1, cpx2)
		}
	}

	var d float64
	d = math.Sqrt(9.0)
	fmt.Printf("%v\n", d)
	d = math.Sqrt(-9.0)
	fmt.Printf("%v\n", d)
}
