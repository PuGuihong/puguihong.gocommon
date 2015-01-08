//calc.go
package main

import (
	"fmt"
	"os"
	"simplemath"
	"simpletest"
	"strconv"
)

var Usage = func() {
	fmt.Println("USAGE: calc command [arguments] ...")
	fmt.Println("\nThe commands are:\n\tadd\tAddition of two values.\n\tsqrt\tSquare root of a non-negative value.")
}

func main() {

	//simpletest.PrintSample()//测试变量
	//simpletest.Sample1()//测试bool
	//simpletest.Sample2() //测试整型数据
	//simpletest.Sample3() //测试浮点数据
	//simpletest.Sample4() //测试复数数据
	//simpletest.Sample5() //测试字串数据
	//simpletest.Sample6() //测试数组数据
	//simpletest.Sample7() //测试数组切片数据
	simpletest.Sample8() //测试map数据
	return
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
