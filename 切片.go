package main

import "fmt"

func qdemo1() {
	// 声明切片类型
	var a []string              //声明一个字符串切片
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	var d = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)              //[]
	fmt.Println(b)              //[]
	fmt.Println(c)              //[false true]
	fmt.Println(a == nil)       //true
	fmt.Println(b == nil)       //false
	fmt.Println(c == nil)       //false
	//fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较
	fmt.Println(nil == d) //切片是引用类型，不支持直接比较，只能和nil比较
}
func qdemno2() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // s := a[low:high]
	// 切片的容量是从切片第一个元素到最后一个元素的数量
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
}
func qdemo3() {
	a := make([]int, 2, 10)
	fmt.Println(a)      //[0 0]
	fmt.Println(len(a)) //2
	fmt.Println(cap(a)) //10
}

//遍历
func qdemo4() {
	s := []int{1, 3, 5}

	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}

	for index, value := range s {
		fmt.Println(index, value)
	}
}

//添加
func qdemo5() {
	var s []int
	s = append(s, 1)       // [1]
	s = append(s, 2, 3, 4) // [1 2 3 4]
	s2 := []int{5, 6, 7}
	s = append(s, s2...) // [1 2 3 4 5 6 7]
	fmt.Println(s)
}

//赋值
func qdemo6() {
	// copy()复制切片
	a := []int{1, 2, 3, 4, 5}
	c := make([]int, 5, 5)
	copy(c, a)     //使用copy()函数将切片a中的元素复制到切片c
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(c) //[1 2 3 4 5]
	c[0] = 1000
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(c) //[1000 2 3 4 5]
}

//要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断。
func main() {
	//qdemo1()
	//qdemno2()
	qdemo6()
}
