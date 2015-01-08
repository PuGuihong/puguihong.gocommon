//filesimple.go
package simpletest

import (
	"os"
)

func FileHelpe() {
	var file = File.Create("E:/sample.txt")
	file.WriteString("test")
	file.Close()
}
