/*
单元通信模式二  - 消息机制通信

1.每个并发单元是自包含的、独立的个体，并且都有自己的变量，但在不同并发单元间
这些变量不共享，每个并发单元的输入和输出只有一种，那就是消息。
2.GO提供的消息通信机制被称为channel。
3.不要通过共享内存来通信，而应该通过通信来共享内存。
4.channel可以在两个或多个goroutine之间传递消息，channel是进程内的通信方式，
通过channel传递对象的过程和调用函数时的参数传递行为比较一致，比如也可以传递指针。
如果需要跨进程通信，建议用分布式系统的方法来解决，比如使用socket或者http等通信协议。
5.channel是类型相关的，一个channel只能传递一种类型的值，这个类型需要在声明channel
时指定。
6.向channel写入数据通常会导致程序阻塞，直到有其他goroutine从这个channel中读取数据。
7.如果channel之前没有写入数据，从channel中读取数据也会导致程序阻塞，直到channel中被写入数据为止。
8.select机制：通过调用select函数来监控一系列的文件句柄，一旦其中一个文件句柄发生了IO动作，该select()
调用就会被返回。后来该机制也被用于实现高并发的Socket服务器程序。GO在语言级别支持select关键字，用于处理异步io问题。
其中最大的一条限制就是case语句里必须是一个io操作。
9.channel缓冲机制(适用于持续传输大量数据的场景)。
10.并发编程错误处理 - 超时机制
错误发生场景：向channel写入数据时发现channel已满，或者充channel师徒读取数据时返现channel为空，如果不正确
处理这些情况，很可能会导致整个goroutine锁死。

Channel 属性和方法
1.channel的传递
2.单向channel:只能用于发送或者接收数据，channel本身必然是同时支持读写的。
单向channel只是对channel的一种使用限制,我们在将一个channel变量传递到一个函数时，通过将其指定为单向channel变量，
从而限制该函数中可以对此channel的操作，比如只能往这个channel写，或者只能从这个channel读取。
3.关闭channel，如何判断channel是否已经被关闭，可以通过使用多重返回值的方式。

多核化并行
1.在执行一些昂贵的计算任务时，我们希望能够尽量利用现代服务器普遍具备的多核特性来尽量将任务并行化，从而达到降低总计算时间的目的，
我们需要了解cpu核心的数量，并针对性地分解计算任务到多个goroutine中去并行运行。

资源同步：
1.go语言包中的sync包提供了两种锁类型：sync.Mutex,sync.RWMutex
2.全局唯一性操作 ：once
*/
package goroutinesimple

import (
	"fmt"
	"sync"
	"time"
)

/*
全局唯一性操作
*/
var a string
var once sync.Once

func setup() {
	a = "hello world"
}
func doprint() {
	once.Do(setup)
}
func twoprint() {
	go doprint()
	go doprint()
}

type Vector []float64

//分配给每个cpu的计算任务
func (v Vector) doSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		v[i] += v[i]
	}
	c <- 1 //发信号高速任务管理者，我已经计算完成了。
}

const NCPU = 2

func (v Vector) doAll(u Vector) {
	c := make(chan int, NCPU)
	for i := 0; i < NCPU; i++ {
		go v.doSome(i*len(v)/NCPU, (i+1)*len(v)/NCPU, u, c)
	}
	//等待所有cpu的任务完成
	for i := 0; i < NCPU; i++ {
		<-c //获取到一个数据，表示一个cpu计算完成了
	}
}

/*
channel 的单向传递
*/
var (
	ch1 chan int       //chan1是正常的channel，不是单向的
	ch2 chan<- float64 //ch2是单向channel，只用于写float64类型数据
	ch3 <-chan int     //ch3是单向channel，只用于读取int数据
)

func parse(ch <-chan int) {
	for value := range ch {
		fmt.Println("Parsing value", value)
	}
}

func Sample7() {
	ch4 := make(chan int)
	ch5 := <-chan int(ch4) //ch5就是一个单向的读取channel
	ch6 := chan<- int(ch4) //ch6是一个单向的写入channel
	parse(ch4)
	fmt.Println(ch5, ch6)
}

/*
channel的传递
*/
type PipeData struct {
	value   int
	handler func(int) int
	next    chan int
}

func Sample6(queue chan *PipeData) {
	for data := range queue {
		data.next <- data.handler(data.value)
	}
}

/*
并发错误处理 - 超时机制
*/
func Sample5() {
	timeout := make(chan bool, 1)
	//创建带1024缓冲大小的channel,即使没有读取方，写入防也可以一直往channel里写入，
	//在缓冲区被填完之前都不会阻塞
	c := make(chan int, 1024)
	for {
		select {
		case c <- 0:
		case c <- 1:
		}
	}
	go func() {
		time.Sleep(1e9) //等待1秒
		timeout <- true
	}()
	select {
	case <-c: //从c中读取数据
		//从带缓冲的channel中读取数据可以适用于常规非缓冲channel完全一致的方法，我们也
		//可以用range关键字来实现更为简便的循环读取：
		for i := range c {
			fmt.Println("Received :", i)
		}
	case <-timeout: //一直没有从c中读取到数据，但是从timeout中读取到了数据
	default:
	}
}

var chName chan int        //声明chan
var m map[string]chan bool //声明chan

//select方式实现异步IO
func Sample4() {
	ch := make(chan int, 1)
	for {
		select {
		case ch <- 0:
		case ch <- 1:
		}
		i := <-ch
		fmt.Println("value received:", i)
	}
}

func Count2(ch chan int) {
	ch <- 1
	fmt.Println("Counting")
}

func Sample3() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count2(chs[i])
	}
	for _, ch := range chs {
		fmt.Println(" ", <-ch, "\n")
	}
}
