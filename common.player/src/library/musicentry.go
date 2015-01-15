//musicentity.go
//音乐信息：唯一ID, 音乐名, 艺术家名 ,音乐位置 ,音乐文件类型
package library

type MusicEntry struct {
	Id     string //唯一ID
	Name   string //音乐名
	Artist string //艺术家名
	Source string //音乐位置
	Type   string //音乐文件类型(MP3,WAV)
}
