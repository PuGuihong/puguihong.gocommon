//simpletest.go
/*
1.不管是什么平台，什么编程语言，不管在哪，并发都是一个大话题，并发的难度在于协调，而协调就要通过交流。
并发单元间的通信是最大的问题。
2.常见并发通信模型：共享数据和消息
共享数据是指多个并发单元分别保持对同一个数据的引用，实现对该数据的共享。被共享的数据可能有多种形式，
比如内存数据块、磁盘文件、网络数据等，在实际工程应用中最常见的无疑是内存。
*/
package goroutinesimple

import (
	"fmt"
	"runtime"
	"sync"
)

/*
单元通信模式一 - 共享内存方式通信
*/
var counter int = 0

func Count(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println("Z")
	lock.Unlock()
}

//单元通信调用
func Sample2() {
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go Count(lock)
	}
	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
}

//加法
func Add(a, b int) int {
	var result = a + b
	fmt.Println("计算结果是：", result, "\n")
	return result
}

//协程并发执行方法
//1.并发方法的初始化和调用。
//2.并发通信
func Sample1() {
	for i := 0; i < 10; i++ {
		go Add(i, i)
	}
}
