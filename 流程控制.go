package main

import (
	"fmt"
)

/**
if 表达式1 {
    分支1
} else if 表达式2 {
    分支2
} else{
    分支3
}
*/
/**
for 初始语句;条件表达式;结束语句{
    循环体语句
}
*/
func ifDemo1() {
	score := 95
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}

// if条件判断特殊写法
func ifDemo2() {
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}
func forDemo() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
func rangDemo() {
	s := "hello济南"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}
}
func switchDemo1() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}
}
func table99() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, j*i)
		}
		fmt.Println()
	}
}
func main() {
	ifDemo1()
	ifDemo2()
	forDemo()
	rangDemo()
	switchDemo1()
	table99()
}
