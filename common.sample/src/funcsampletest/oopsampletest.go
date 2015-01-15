//oopsimpletest.go
//面向对象的编程方式
package funcsampletest

import (
	"fmt"
)

//成员变量的可访问性测试
func OopSample7() {
	r := &MyRect{X: 12.0, Y: 13.2, Width: 21.31, Height: 32.12}
	area := r.area()
	fmt.Println("矩形类 成员变量是： X = ", r.X, " ,Y = ", r.Y, " ,Width = ", r.Width, " ,Height = ", r.Height, "\n")
	fmt.Println("矩形的成员方法获取面积是：", area)
}

//Rect的公有成员变量与方法
//GO语言的可访问性是包一级的而不是类一级的。
type MyRect struct {
	X, Y          float64
	Width, Height float64
}

func (r *MyRect) area() float64 {
	return r.Width * r.Height
}

// 组合应用
func OopSample6() {
	bf1 := new(base)
	bf1.fbase()
	ff1 := new(foo)
	ff1.ffoo()
}

///匿名组合
type base struct {
	name string
}

//实现基础类型
func (base *base) fbase() {
	base.name = "Foo"
	fmt.Println(base.name)
}

type foo struct {
	base
}

func (foo *foo) ffoo() {
	foo.base.name = "Foo"
	fmt.Println(foo.base.name)
}

//实现Rect自定义类型
//GO语言中没有构造函数的概念，对象的创建有全局创建函数来创建完成，以NewRect来表示“构造函数”
//Rect的“构造函数”
//func NewRect(x, y, width, height float64) *Rect {
//	return &Rect(x, y, width, height)
//}
func OopSample4() {
	//创建和初始化自定义Rect类型的方法
	rect1 := new(Rect) //未进行显式初始化的变量都会被初始化为该类型的零值
	rect2 := &Rect{}
	rect3 := &Rect{0, 0, 100, 200}
	rect4 := &Rect{width: 100, height: 200}
	fmt.Println("打印自定义类型执行方法的结果：", rect1.Area(), rect2.Area(), rect3.Area(), rect4.Area())
}

//结构体 ,和class地位同等，但是只有composition基础特性
//所有的go语言类型都可以有自己的方法（指针类型除外）
//自定义Rect矩形类型
type Rect struct {
	x, y          float64
	width, height float64
}

//定义Rect类型的实现方法
func (r *Rect) Area() float64 {
	return r.width * r.height
}

//特别的 值语义和引用语义
//1.数组切片 指向数组的一个区间;
//2.map 常见数据结构，提供键值对查询能力;
//3.channel 执行体goroutine间通信的设施;
//4.interface 对一组满足某个契约的类型的抽象
func OopSample3() {

}

type Integer int

//值语义和引用语义
//a,b 两个值， b=a b.Modify();如果b的修改不会影响到a的值，那么此类型为值类型，否则是引用类型
func OopSample2() {
	//数组是值类型，b=a赋值后会将内容进行完整的复制
	var (
		a = [3]int{1, 2, 3}
		b = a
		c = &a
	)
	b[1]++
	fmt.Println("b = a方式赋值后:", a, b, " \n")
	//如果想要改变值类型的值，必须要用到指针，表达引用

	c[1]++
	fmt.Println("c = &a 方式赋值后:", a, b, " \n")
}

//面向对象方式，实现为自定义Integer 类型添加方法
func (a Integer) Less(b Integer) bool {
	return a < b
}

//面向过程方式，实现自定义Integer 类型添加自定义方法
func Integer_Less(a Integer, b Integer) bool {
	return a < b
}

//用到指针，修改自定义Integer对象时
//适用于：传入的对象很大的时候，如果传入对象过小(4 bit)，用指针会产生额外的开销，并不划算
func (a *Integer) Add(b Integer) {
	*a += b
}

//用到值，不能修改自定义Integer对象的值，因为传递的是值类型的数据
func (a Integer) Add2(b Integer) {
	a += b
}
func OopSample1() {
	//面向对象的调用方式
	var (
		a Integer = 1
		b Integer = 2
	)

	if a.Less(2) {
		fmt.Println(a, "Less 2")
	}
	//面向过程的调用方式
	if Integer_Less(a, b) {
		fmt.Println(a, "Less b:", b)
	}
	//用指针对象进行类型的自我修改，需要修改变量的值时只能传递指针
	a.Add(2)
	fmt.Println("a = ", a)
	//用值进行进行传递时，不能实现对象类型的自我修改
	b.Add2(1)
	fmt.Println("b = ", b)
}
