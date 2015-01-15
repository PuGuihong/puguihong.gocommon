//basicsimple.go
//1.类的继承树没有实际意义，在接口应用中，只需要知道这个类实现了哪些方法，每个方法是什么含义。
//2.实现类的时候，只需要关心自己应该提供哪些方法，不用纠结接口需要拆的多细才合理，接口由使用方按需要定义，不用事前规划。
//3.不用为了实现一个接口而导入一个包，接口由使用方按自身需求来定义，使用方无需关心是否有其他模块定义过赖斯的接口。
package interfacesimpletest

import (
	"fmt"
)

//2.将一个接口赋值给另一个接口
//只要两个接口拥有相同的方法列表（不论次序），两个接口就是等同的，可以互相赋值
type readwrite interface {
	read(buf []byte) (n int, err error)
	write(buf []byte) (n int, err error)
}
type istream interface {
	write(buf []byte) (n int, err error)
	read(buf []byte) (n int, err error)
}
type write interface {
	write(buf []byte) (n int, err error)
}
type file struct {
}

func (f *file) read(buf []byte) (n int, err error) {
	var errr error
	return 1, errr
}

func (f *file) write(buf []byte) (n int, err error) {
	var errr error
	return 1, errr
}

//接口赋值并不要求两个接口必须等价，如果A的方法列表是接口B的方法列表的子集，B可以赋值给接口A
var (
	lfile1 istream   = new(file)
	lfile2 readwrite = lfile1
	lfile3 istream   = lfile2
	lfile4 write     = lfile1
	//lfile4 istream = lfile1//不成立，lfile1 没有read()方法
)

//接口查询
//接口查询是否成功，要在运行期才能够确定；不像接口赋值，编译器只需要静态类型检查即可以判断赋值是否可行。
//GO内置了接口查询功能。
func Sample2() {
	//	var lfile1 istream   = new(file)
	//	var lfile5 write   = lfile1
	//检查lfile1指向的对象实例是否实现了istream接口，如果实现了，执行特定的代码
	//	if lfile5, ok := lfile5.(istream){}
}

//1.将对象实例赋值给接口：要求该对象实例实现了接口要求的所有方法
type integer int

func (a integer) less(b integer) bool {
	return a < b
}
func (a *integer) add(b integer) {
	*a += b
}

type lessAdder interface {
	less(b integer) bool
	add(b integer)
}
type lesser interface {
	less(b integer) bool
}

//接口赋值
var (
	a integer   = 1
	b lessAdder = &a
	//b  lessAdder = a//考虑到add方法
	b2 lesser = &a
	b3 lesser = a
)

//定义接口  用于File类的继承
type IFile interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	Seek(off int64, whence int) (pos int64, err error)
	Close() error
}
type IReader interface {
	Read(buf []byte) (n int, err error)
}
type IWriter interface {
	Write(buf []byte) (n int, err error)
}
type ICloser interface {
	Close() error
}

//File类并没有从这些接口继承，甚至不知道这些接口的存在，但File类实现了这些接口，可以进行赋值。
//对象赋值的两种情况：
//1.将对象实例赋值给接口：要求该对象实例实现了接口要求的所有方法
//2.将一个接口赋值给另一个接口。
var (
	file1 IFile   = new(File)
	file2 IReader = new(File)
	file3 IWriter = new(File)
	file4 ICloser = new(File)
)

//自定义类型 方式实现 Read 、Write、Seek、Close方法
type File struct {
}

func (f *File) Read(buf []byte) (n int, err error) {
	var errr error
	return 1, errr
}

func (f *File) Write(buf []byte) (n int, err error) {
	var errr error
	return 1, errr
}
func (f *File) Seek(off int64, whence int) (pos int64, err error) {
	var errr error
	return 1, errr
}
func (f *File) Close() error {
	var err error
	return err
}

func Sample1() {
	fmt.Println("tst")
}
