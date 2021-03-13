package main

import (
	"fmt"
	"math"
	"strings"
)

// 修改字符串
func changeString() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}

// 类型转换
func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a) // 10
	fmt.Printf("%b \n", a) // 1010  占位符%b表示二进制

	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b) // 77

	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) // ff
	fmt.Printf("%X \n", c) // FF

	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)

	fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")
	s1 := `第一行
第二行
第三行`
	fmt.Printf("%T \n", s1)
	fmt.Printf("%v \n", s1)
	fmt.Printf("%#v \n", s1)
	// 分隔
	ret := strings.Split(s1, "\n")
	fmt.Println(ret)
	ss := "hello world济南"
	n := len(ss)
	for i := 0; i < n; i++ {
		//fmt.Println(ss[i])
		fmt.Printf("%c\n", ss[i])
	}
	for _, c := range ss {
		fmt.Printf("%c \n", c)
	}
	changeString()
	sqrtDemo()
}
