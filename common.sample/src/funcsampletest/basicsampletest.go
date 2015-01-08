//basicsampletest.go
package funcsampletest

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func Sample() {
	fmt.Println("test")
}

func Add(a int, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("需要两个非负数")
		return
	}
	return a + b, nil //支持多重返回值
}

func Add2(a, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("需要两个非负数")
		return
	}
	return a + b, nil //支持多重返回值
}

func Add3(a, b, c int) (ret, num int) {
	return a + b + c, 3
}

func Sample2() {
	myfun(2, 1, 3, 4)
}

//不定参数函数
func myfun(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func Sample3() {
	var (
		v1 int     = 1
		v2 float64 = 53.2
		v3 string  = "hello"
		v4 float32 = 1.234
	)
	myPrintf(v1, v2, v3, v4)
}

//不定参数 interface
func myPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, " 是 int 值")
		case string:
			fmt.Println(arg, " 是 string 值")
		case int64:
			fmt.Println(arg, " 是 int64 值")
		default:
			fmt.Println(arg, "是 未知类型的值")
		}
	}
}
func Sample4() {
	copyFile("E:/download_14_12_30_12_40_14.jpg", "D:/2.jpg")
}

//多返回值
//defer 方法
func copyFile(dst, src string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}
	defer srcFile.Close()
	dstFile, err := os.Create(dst)
	if err != nil {
		return
	}
	defer dstFile.Close()
	return io.Copy(dstFile, srcFile)
}

//panic() , recover() 内置错误处理函数
/*
当在一个函数执行过程中调用panic()函数时，正常的函数执行流程将立即终止，但函数中 之前使用defer关键字延迟执行的语句将正常展开执行，之后该函数将返回到调用函数，并导致 逐层向上执行panic流程，直至所属的goroutine中所有正在执行的函数被终止。错误信息将被报 告，包括在调用panic()函数时传入的参数，这个过程称为错误处理流程。
recover()函数用于终止错误处理流程。一般情况下，recover()应该在一个使用defer 关键字的函数中执行以有效截取错误处理流程。如果没有在发生异常的goroutine中明确调用恢复 过程（使用recover关键字），会导致该goroutine所属的进程打印异常信息后直接退出。
*/
