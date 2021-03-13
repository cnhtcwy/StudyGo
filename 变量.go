package main

import "fmt"

// 声明变量
var name string
var age int
var isOk bool

//批量声明
var (
	name1 string
	age1  int
	noOk  bool
)

//s5 := "呵呵，不能用"
func foo1() (int, string) {
	return 10, "aaa"
}

func main() {
	name = "wangyang"
	age = 18
	isOk = true
	age1 = 20
	fmt.Printf("name:%s\n", name)
	fmt.Printf("age:%d", age)
	var s1 string = "acme"
	print(s1)
	var s2 = "admin"
	print(s2)
	print("\n")
	// 只能用于局部变量
	s3 := "sa"

	print(s3)
	fmt.Print("\n")

	x, _ := foo1()
	_, y := foo1()
	fmt.Println("x=", x)
	fmt.Println("y=", y)
}
