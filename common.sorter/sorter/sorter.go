//sorter.go 排序算法比较成都DEMO,主程序部分
//功能：
//1.从命令行指定输入的数据文件和输出的数据文件，
//2.指定对应的排序算法
//用法示例：USAGE: sorter -i <in> -o <out> -a <qsort|bubblesort>
//$./sorter -I in.dat -o out.dat -a qsort
//输入不合法给出提示

//主程序功能：
//1.获取并解析命令行输入；
//2.从对应文件夹中读取输入数据
//3.调用对应的排序函数
//4.将排序的结果输入到对应的文件夹中。
//5.打印排序所花费的时间信息。

//包说明：
//flag		快速解析命令行参数
//os		系统文件操作
//bufio 	文件读写
//io		文件读写
//strconv	字符串处理
package main

import (
	"algorithms/bubblesort"
	"algorithms/qsort"
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

var infile *string = flag.String("i", "infile", "File	contains	values	for	sorting")
var outfile *string = flag.String("0", "outfile", "File	to	receive	sorted	values")
var algorithm *string = flag.String("a", "qsort", "Sort	algorithm")

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create the output file", outfile)
		return err
	}
	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open the input file", infile)
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			fmt.Println("A too long line,seems unexpected.")
			return
		}

		str := string(line) //转换字符数组为字符串

		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}
	return
}

func main() {
	flag.Parse()
	if infile != nil {
		fmt.Println("infle =", *infile, " outfile =", *outfile, "algorithm =", *algorithm)
	}

	values, err := readValues(*infile)
	if err == nil {
		t1 := time.Now()
		switch *algorithm {
		case "qsort":
			qsort.QuickSort(values)
		case "bubblesort":
			bubblesort.BubbleSort(values)
		default:
			fmt.Println("Sorting algorithm ", *algorithm, "is either unknown or unsuppotred.")
		}
		t2 := time.Now()
		fmt.Println("The sorting process costs ", t2.Sub(t1), " to complete")
		fmt.Println("Read values:", values)
		writeValues(values, *outfile)
	} else {
		fmt.Println(err)
	}
}
