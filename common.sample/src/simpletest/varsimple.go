//varsimple.go
/*
*GO 基础内置类型
布尔 bool
整型 int8,byte,int16,int,uint,uintptr
浮点类型 float32,float64
复数类型 complex64,complex128
字串 string
字符类型 rune
错误类型 error
**GO 复合类型
指针 pointer
数组 array
切片 slice
字典 map
通道 chan
结构体 struct
接口 interface
*
*/
package simpletest

import (
	"fmt"
)

//自定义常量
const (
	Pi      float64 = 3.14159265358979323846
	zero            = 0.0
	u, v    float32 = 0, 3
	a, b, c         = 2, 3, "foo"
)

//定义枚举
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numberofDays
)

//自定义变量
var (
	v1 int
	v2 string
	//v3 [10]int
	//v4 []int
	/*v5 struct {
		f int
	}
	v6 *int
	v7 map[string]int
	v8 func(a int) int*/
)

//自定义变量，自定义常量，预定义常量，枚举
func PrintSample() {
	v1 = 10
	v2 = "This is string"
	//v3 = {1,2,3,4,5,6,7,8,9,0}
	//v4 = new []int{[1]}

	fmt.Println("自定义变量的值：", v1, v2)
	fmt.Println("自定义常量的值：", Pi, zero, u, v, a, b, c)
}

//bool 类型数据
func Sample1() {
	var v1 bool
	v1 = true
	//支持类型的推导,不支持强制赋值和强制转换
	v2 := (1 == 2)
	fmt.Println(" bool类型数据：", v1, v2)
}

//整型 类型数据
func Sample2() {
	var val2 int32
	val1 := 64
	//int int32必须进行强制转换,转换后会损失精度
	val2 = int32(val1)
	//整型数据的 运算关系 + ,- ,*, /, %取余数
	val3 := (5 % 3)
	fmt.Println("整型 类型数据:", val1, val2, val3)

	//整型数据的 比较关系 > ,< ,== ,>= ,<= ,!= ; 整型类型变量不能直接进行比较，但是可以和字面常量比较
	if val1 == 1 || val2 == 2 {
		fmt.Println("可以进行字面常量的直接比较，但是不能 val1 == val2 ")
	}
	//整型数据类型的位运算 x<<y(左移)，x>>y(右移)，x^y(亦或)，x&y(与)，x|y(或)，^x(取反)
	fmt.Println("整型数据类型的位运算：", 124<<2, 124>>2, 124^2, 124&2, 124|2, ^2)
}
