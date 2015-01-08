//filesimple.go
package simpletest

import (
	"os"
)

func FileHelpe() {
	file, err := os.Create("E:/sample2.txt")
	if err != nil {
		file.WriteString("this is test file")
	}
	file.Close()
}
