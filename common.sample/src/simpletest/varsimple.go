//varsimple.go
package simpletest

import (
	"fmt"
)

var (
	v1 int
	v2 string
	//v3 [10]int
	//v4 []int
	/*v5 struct {
		f int
	}
	v6 *int
	v7 map[string]int
	v8 func(a int) int*/
)

func PrintSample() {
	v1 = 10
	v2 = "This is string"
	//v3 = {1,2,3,4,5,6,7,8,9,0}
	//v4 = new []int{[1]}

	fmt.Println(v1, v2)
}
