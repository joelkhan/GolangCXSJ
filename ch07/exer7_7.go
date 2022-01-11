package main

import (
	"bytes"
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("函数defer开始运行……")
		if err := recover(); err != nil {
			fmt.Println("程序异常退出：", err)
		} else {
			fmt.Println("程序正常退出")
		}
	}()

	/*
	 * 按照题目要求，准备的测试数据
	 */
	//username := "joel"
	//username := "Mike D."
	//username := "jim,around"
	// username := "#Nickle's#Back#"
	username := "ABCisGood!"
	f(username)
}

func f(user string) {
	fmt.Println(user)
	fmt.Println("函数f开始运行……")
	if checkInvalidName(user) {
		panic("username包含非法字符！")
	} else {
		fmt.Println("函数f调用结束。")
	}
}

func checkInvalidName(name string) (res bool) {
	nameSlice := []byte(name)
	if bytes.IndexFunc(nameSlice, checkInvalidF) == -1 {
		res = false
	} else {
		res = true
	}
	return res
}

func checkInvalidF(r rune) (res bool) {
	if r == ' ' || r == ',' || r == '#' || r == '!' {
		res = true
	} else {
		res = false
	}
	return
}
