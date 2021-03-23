package main

import "fmt"

type dog struct{}

type cat struct{}

// dog实现了Sayer接口
func (d dog) say() {
	fmt.Println("汪汪汪")
}

// cat实现了Sayer接口
func (c cat) say() {
	fmt.Println("喵喵喵")
}
func (d *dog) move() {
	fmt.Println("狗会动")
}

//var x Sayer // 声明一个Sayer类型的变量x
func main() {

	a := cat{} // 实例化一个cat
	b := dog{} // 实例化一个dog
	d := &dog{}
	a.say()
	b.say()
	d.move()
	//x = a       // 可以把cat实例直接赋值给x
	//x.say()     // 喵喵喵
	//x = b       // 可以把dog实例直接赋值给x
	//x.say()     // 汪汪汪
}
