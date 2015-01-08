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
	"math"
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

//浮点 类型数据
func Sample3() {
	var fval1 float64 = 3.1415926535897932385
	fval2 := 12.0              //go 会自定推导为浮点，并且自动设为float64
	var fval4 = float32(fval2) //必须进行强制转换
	var fval3 float64 = 3.14169265358989323846
	p := 0.0000001 //比较精度
	//浮点数比较不能直接用 == 比较 ，需要math.Fdim进行精确的比较,表示在设置精度范围内两浮点数值相等
	if math.Dim(fval1, fval3) < p {
		fmt.Println("在精度范围内，两浮点数相等")
	} else {
		fmt.Println("在精度范围内，两浮点数不相等")
	}
	fmt.Println("浮点型数据：", fval1, fval2, fval3, fval4)
}

//复数类型数据
func Sample4() {
	var val1 complex64 //由两个float32构成的复数数据类型
	val1 = 3.2 + 12i
	val2 := 3.2 + 12i        //complex128类型的复数
	val3 := complex(3.2, 12) //val3 结果同val2
	fmt.Println("复数 数据：", val1, real(val1), imag(val1), val2, real(val2), imag(val2), val3, real(val3), imag(val3))
}

//字串 数据类型
func Sample5() {
	var str string
	str = "Hello , 你好世界 ll"
	fmt.Printf("字串 \"%s\"的长度是: %d \n", str, len(str))
	//字串遍历，字符串的内容可以用类似于数组下标的方式获取，但与数组不同，字符串的内容不能在初始 化后被修改
	for i, ch := range str {
		fmt.Printf(" 第%d个字符是%c \n", i, ch)
	}
	//以Unicode字符方式遍历时，每个字符的类型是rune
	for i, ch := range str {
		fmt.Println(i, ch)
	}
	//字串遍历2，根据下标提取字串中的字符，获取类型为byte
	n := len(str)
	for i := 0; i < n; i++ {
		ch := str[i]
		fmt.Println(i, ch)
	}
}

//array数组 数据类型
//1.数组长度在定义后就不可更改,可以由len()获取数组长度，进行下标遍历
//2.可以通过内置的rangeg关键字进行遍历
//3.数组是值类型，所有值类型变量在赋值和作为参数传递时都将产生一次赋值动作；数组作为函数参数与后会被复制，函数体无法修改传入的数组内容，只是在操作传入数组的副本。
func Sample6() {
	var (
		arra1 [32]byte                    //长度为32的字节数组
		arra2 [2 * 2]struct{ x, y int32 } //复杂类型数组
		arra3 [1000]*float64              //指针数组
		arra4 [3][5]int                   //二维数组存储3*5个数据
		arra5 [2][2][2]float64            //等同于[2]([2]([2]float64))
	)
	//对数组进行下标的遍历
	for i := 0; i < len(arra1); i++ {
		fmt.Println("Element ", i, " of array is ", arra1[i])
	}
	//range进行遍历
	for i, v := range arra1 {
		fmt.Println(i, v)
	}
	fmt.Println("\narra1 的长度：", len(arra1), "\narra2 的长度：", len(arra2), "\narra3 的长度：", len(arra3), "\narra4 的长度：", len(arra4), "\narra5 的长度：", len(arra5))
}

//slice数组切片 数据类型;类似一个指向数组的指针
//抽象解释：一个指向原生数组的指针、数组切片中的元素个数、数组切片已分配的存储空间
func Sample7() {
	//创建数组切片方法一，基于数组创建切片
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var mySlice []int = myArray[:5]
	var mySlice2 []int = myArray[:]
	var mySlice3 []int = myArray[5:]
	fmt.Println("数组中的元素是：\n")
	for _, v := range myArray {
		fmt.Println(v, " ")
	}
	fmt.Println("\nmySlice切片中的数据是：")
	for _, v := range mySlice {
		fmt.Println(v, " ")
	}
	fmt.Println("\nmySlice2切片中的数据是：")
	for _, v := range mySlice2 {
		fmt.Println(v, " ")
	}
	fmt.Println("\nmySlice3切片中的数据是：")
	for _, v := range mySlice3 {
		fmt.Println(v, " ")
	}
	//创建数组切片方法二,利用make()方法直接创建切片
	mySlice4 := make([]int, 5)       //创建初始元素个数为5的int数组切片,元素初始值为0
	mySlice5 := make([]int, 5, 10)   //创建初始元素个数为5的数组切片，元素初始值为0，并预留10个元素的存储空间
	mySlice6 := []int{1, 2, 3, 4, 5} //直接创建并且初始化5个元素的数组切片

	//切片的操作性一 ：动态增减元素,元素的个数和分配的空间可以是两个不同的概念
	//合理设置切片存储能力值，可以降低切片重新分配内存和操作内存频率提供性能
	//cap()函数返回切片分配空间的大小，len()返回数组切片中当前所存储的元素个数
	fmt.Println("存储空间的大小：", cap(mySlice4), cap(mySlice5), cap(mySlice6))
	fmt.Println(len(mySlice4), len(mySlice5), len(mySlice6))
	//append()新增切片元素
	mySlice4 = append(mySlice4, 1, 2, 3)
	mySlice5 = append(mySlice4, mySlice6...)
	fmt.Println("存储空间的大小：", cap(mySlice4), cap(mySlice5), cap(mySlice6))
	fmt.Println(len(mySlice4), len(mySlice5), len(mySlice6))

	//创建数组切片三 ，基于切片创建切片
	mySlice7 := mySlice6[:4]
	fmt.Println(cap(mySlice7), len(mySlice7))
	//copy()切片数组的内容复制
	mySlice8 := []int{1, 2, 3, 4, 5}
	mySlice9 := []int{5, 4, 3}
	copy(mySlice8, mySlice9) //只会复制mySlice8的前三个元素到mySlice9中
	fmt.Println("\nmySlice8切片中的数据是：")
	for _, v := range mySlice8 {
		fmt.Println(v, " ")
	}
	fmt.Println("\nmySlice9切片中的数据是：")
	for _, v := range mySlice9 {
		fmt.Println(v, " ")
	}
	copy(mySlice9, mySlice8) //只会复制mySlice9中的三个元素到mySlice8的前三个位置
	fmt.Println("\nmySlice8切片中的数据是：")
	for _, v := range mySlice8 {
		fmt.Println(v, " ")
	}
	fmt.Println("\nmySlice9切片中的数据是：")
	for _, v := range mySlice9 {
		fmt.Println(v, " ")
	}
}

//结构体
type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

//map 图 数据类型
//map是一堆键值对的未排序集合
//1.变量声明：var myMap map[string] PersonInfo myMap是声明的map变量名，string是键的类型，PersonInfo则是其中所存放的值类型
//2.创建: myMap = make(map[string]PersonInfo) ; myMap = make(map[string] PersonInfo,100)
//3.直接创建并且初始化：myMap = map[string]PersonInfo{"123":PersonInfo{"1","Jack","room 111"}
//4.元素赋值：myMap["1234"] = PersonInfo{"1","Jack","romm 111"}
//5.元素删除:delete(myMap,"1234")
//6.元素查找：value,ok := myMap["1234"] if ok{}
func Sample8() {
	//身份证号作为唯一标示
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)

	personDB["123"] = PersonInfo{"123", "owen pu", "room 212"}
	personDB["124"] = PersonInfo{"123", "owen yan", "room 215"}

	//查找map中的数据
	person, ok := personDB["123"]
	if ok {
		fmt.Println("找到元素", person.Name, "id 表示 123")
	} else {
		fmt.Println("没有找到数据")
	}
}
