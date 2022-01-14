package main

import (
	"fmt"
)

// 利用结构体定义对象：People、Teacher、Student
type People struct {
	Name string
}
type Teacher struct {
	People
	Department string
}
type Student struct {
	People
	School string
}

// 对象方法实现
func (p People) SayHi() {
	fmt.Printf("Hi, I'm %s. Nice to meet you!\n", p.Name)
}

func (t Teacher) SayHi() {
	fmt.Printf("Hi, my name is %s. I'm working in %s.\n", t.Name, t.Department)
}

func (s Student) SayHi() {
	fmt.Printf("Hi, my name is %s. I'm studing in %s.\n", s.Name, s.School)
}

func (s Student) Study() {
	fmt.Printf("I'm learning Golang in %s.\n", s.School)
}

// 定义接口
type Speaker interface {
	SayHi()
}

type Learner interface {
	SayHi()
	Study()
}

func main() {
	people := People{"张三"}
	teacher := Teacher{People{"郑智"}, "Computer science"}
	student := Student{People{"李明"}, "Yale University"}
	var is Speaker
	is = people
	is.SayHi()
	is = teacher
	is.SayHi()
	is = student
	is.SayHi()
	var il Learner
	il = student
	il.Study()

	ifSlice()
}

/* 一个由接口组成的切片 */
func ifSlice() {
	people := People{"李四"}
	teacher := Teacher{People{"赵六"}, "Computer222"}
	student := Student{People{"王二"}, "Yale666"}
	ix := make([]Speaker, 3)
	ix[0], ix[1], ix[2] = people, teacher, student
	for _, v := range ix {
		v.SayHi()
	}
}
