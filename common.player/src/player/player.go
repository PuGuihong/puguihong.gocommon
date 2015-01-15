//mplayer.go
//音乐播放器 功能流程
//1.音乐库功能，使用者可以查看、添加、删除里面的音乐曲目；
//2.播放音乐
//3.支持MP3和WAV，但是也能随时扩展以便支持更多的音乐类型；
//4.退出程序
//音乐播放器 操作命令
//音乐库管理命令：lib ,包括list/add/remove命令
//播放管理：play命令，play后带歌曲名参数
//退出程序： q命令
//20150115-遗留问题：
//1.多任务；目前同时只能执行一个任务，例如音乐正在播放时，用户不能做其他任何操作，音乐播放过程也不应该导致用户界面无法响应，因此播放应该在一个单独的县城中，并且能够与主程序相互通信。
//  像一般播放器一样，在播放音乐的同时也需要支持一些视觉效果播放，即至少需要：用户界面线程，音乐播放线程，视频播放线程。
//2.控制播放；当前设计是单任务的，所以播放过程无法接受外部的输入，作为一个成熟的播放器，我们至少需要支持暂停和停止等功能，甚至包括设置当前播放位置等。假设我们已经将播放过程放到一个独立
//  的goroutine中，可以使用GO的channel来实现。
package main

import (
	"bufio"
	"fmt"
	"library"
	"os"
	"strconv"
	"strings"
)

var (
	lib *library.MusicManager
	id  int = 1
)
var ctr1, signal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
		}
	case "add":
		if len(tokens) == 6 {
			id++
			lib.Add(&library.MusicEntry{strconv.Itoa(id), tokens[2], tokens[3], tokens[4], tokens[5]})
		} else {
			fmt.Println("USAGE:lib add <name><artist><source><type>")
		}
	case "remove":
		if len(tokens) == 3 {
			lib.Remove(2)
		} else {
			fmt.Println("USAGE: lib remove <name>")
		}
	default:
		fmt.Println("Unrecognized lib command:", tokens[1])
	}
}

func handlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("USAGE: play <name>")
		return
	}
	e := lib.Find(tokens[1])
	if e == nil {
		fmt.Println("The music", tokens[1], "does not exits.")
		return
	}
	library.Play(e.Source, e.Type)
}

func main() {
	fmt.Println("Enter following commands to control the player: \nlib list -- View the existing music lib \nlib add <name><artist><source><type> -- Add a music to the music lib\nlib remove <name> -- Remove the specified music from the lib\nplay <name> -- Play the specified music \n")
	lib = library.NewMusicManager()
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter command ->")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)
		if line == "q" || line == "e" {
			break
		}
		tokens := strings.Split(line, " ")
		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			handlePlayCommand(tokens)
		} else {
			fmt.Println("Unrecognized command:", tokens[0])
		}
	}
}
