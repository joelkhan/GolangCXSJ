package main

import (
	"book.GolangCXSJ/linksmp"
)

func main() {
	var head *linksmp.Node
	stu1 := linksmp.Node{linksmp.Student{100, "李明"}, nil}
	stu2 := linksmp.Node{linksmp.Student{101, "张晓"}, nil}
	stu3 := linksmp.Node{linksmp.Student{102, "赵琼"}, nil}
	stu4 := linksmp.Node{linksmp.Student{103, "王乐"}, nil}
	stu5 := linksmp.Node{linksmp.Student{99, "老刘"}, nil}
	stu6 := linksmp.Node{linksmp.Student{109, "宋二"}, nil}
	stu7 := linksmp.Node{linksmp.Student{104, "鵬儿"}, nil}
	head = head.Creat()
	head = stu1.Insert(head)
	head = stu2.Insert(head)
	head = stu3.Insert(head)
	head = stu4.Insert(head)
	head = stu5.Insert(head)
	head = stu6.Insert(head)
	head = stu7.Insert(head)
	head.PrintLink()

	head = stu3.Delete(head)
	head.PrintLink()
}
