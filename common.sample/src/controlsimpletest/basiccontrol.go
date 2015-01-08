//basiccontrol.go
package controlsimpletest

import (
	"fmt"
)

//switch case
func Sample1() {
	Num := 8
	switch {
	case 0 <= Num && Num <= 3:

	}
	fmt.Println("hello ")
}

//for 无限循环
func Sample2() {
	sum := 0
	for {
		sum++
		fmt.Println(sum)
		if sum > 100 {
			break
		}
	}
}

//for 多重赋值
func Sample3() {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("a 的原始数据：\n")
	for i, v := range a {
		fmt.Println(i, v)
	}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println("\na 的改变数据：\n")
	for i, v := range a {
		fmt.Println(i, v)
	}
}

//for continue break 控制
func Sample4() {
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 5 {
				break
			}
			fmt.Println(i)
		}
	}
}

//goto
func Sample5() {
	i := 0
HERE:
	fmt.Println(i)
	i++
	if i < 10 {
		goto HERE
	}
}
