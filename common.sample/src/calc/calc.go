//calc.go
package main

import (
	"fmt"
	"netsimpletest"
	"os"
	"simplemath"
	"strconv"
)

var Usage = func() {
	fmt.Println("USAGE: calc command [arguments] ...")
	fmt.Println("\nThe commands are:\n\tadd\tAddition of two values.\n\tsqrt\tSquare root of a non-negative value.")
}

func main() {
	/*基本数据类型 	simpletest*/
	//filesimpletest.PrintSample()//测试变量
	//varsimpletest.Sample1()//测试bool
	//varsimpletest.Sample2() //测试整型数据
	//varsimpletest.Sample3() //测试浮点数据
	//varsimpletest.Sample4() //测试复数数据
	//varsimpletest.Sample5() //测试字串数据
	//varsimpletest.Sample6() //测试数组数据
	//varsimpletest.Sample7() //测试数组切片数据
	//simpletest.Sample8() //测试map数据
	/*基本流程控制  controlsimpletest*/
	//controlsimpletest.Sample1() //switch case 流程测试
	//controlsimpletest.Sample2() //for 流程测试
	//controlsimpletest.Sample3() //for 流程测试
	//controlsimpletest.Sample5() //goto 流程测试
	/*基本函数方法  funcsampletest*/
	//funcsampletest.Sample2() //传入不定参数
	//funcsampletest.Sample3() //传入接口类型的不定参数
	//funcsampletest.Sample4() //defer类型的不定参数
	//funcsampletest.OopSample1() //给系统类型添加方法
	//funcsampletest.OopSample2() //值语义和引用语义
	//funcsampletest.OopSample4() //值语义和引用语义
	//funcsampletest.OopSample6() //值语义和引用语义
	//funcsampletest.OopSample7() //成员变量可访问性测试
	/* 接口属性和功能测试 interfacesimpletest*/
	//interfacesimpletest.Sample1() //
	/*并发和协程测试*/
	//goroutinesimple.Sample1()
	//goroutinesimple.Sample2() //共享内存实现进程间通信
	//goroutinesimple.Sample3() //channel方式实现进程间通信
	//goroutinesimple.Sample4() //select 实现异步io ，死循环
	//goroutinesimple.Sample5()
	args := os.Args
	fmt.Println(" ", args[0], len(args))
	if args == nil || len(args) < 2 {
		Usage()
		return
	}

	switch args[0] {
	case "add":
		if len(args) != 3 {
			fmt.Println("USAGE: calc add <integer1><integer2>")
			return
		}
		v1, err1 := strconv.Atoi(args[1])
		v2, err2 := strconv.Atoi(args[2])
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE: calc add <integer1><integer2>")
			return
		}
		ret := simplemath.Add(v1, v2)
		fmt.Println("Result: ", ret)
	case "sqrt":
		if len(args) != 2 {
			fmt.Println("USAGE: calc sqrt <integer>")
			return
		}
		v, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("USAGE: calc sqrt <integer>")
			return
		}
		ret := simplemath.Sqrt(v)
		fmt.Println("Result: ", ret)
	default:
		Usage()
	}
}
