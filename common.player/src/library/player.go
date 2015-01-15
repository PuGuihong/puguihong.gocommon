//player.go
//播放函数
package library

import (
	"fmt"
	"mediaplay"
)

type Player interface {
	Play(source string)
}

func Play(source, mtype string) {
	var p Player
	switch mtype {
	case "MP3":
		p = &mediaplay.MP3Player{}
	case "WAV":
		p = &mediaplay.WAVPlayer{}
	default:
		fmt.Println("Unsupported music type", mtype)
		return
	}
	p.Play(source)
}
