package main

import "fmt"

const pi = 3.1415926
const e = 2.7182
const (
	Pi = 3.15
	E  = 2.718
)
const (
	n1 = 100
	n2
	n3
)

// iota类似枚举
const (
	a1 = iota
	a2 = 100
	a3 = iota
)
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println(pi)
	fmt.Println(Pi)
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(KB)
	fmt.Println(MB)
}
