package main

import "fmt"

//func 函数名(参数)(返回值){
//	函数体
//}
func f1(x int, y int) int {
	return x + y
}

// 简写
func intSum(x, y int) int {
	return x + y
}
func intSum2(x ...int) int {
	fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}
func intSum3(x int, y ...int) int {
	fmt.Println(x, y)
	sum := x
	for _, v := range y {
		sum = sum + v
	}
	return sum
}
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}
func calc1(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

//定义全局变量num
var num int64 = 10

func testGlobalVar() {
	fmt.Printf("num=%d\n", num) //函数中可以访问全局变量num
}
func testNum() {
	num := 100
	fmt.Printf("num=%d\n", num) // 函数中优先使用局部变量
}
func fn(n int64) int64 {
	if n <= 1 {
		return 1
	}
	return n * fn(n-1)
}
func main() {
	//a := f1(1,2)
	//fmt.Println(a)
	//sum := intSum2(1,2,3,4)
	//s1,s2 := calc(15,10)
	//fmt.Println(s1,s2)
	//testNum()
	abc := fn(3)
	fmt.Println(abc)
}
