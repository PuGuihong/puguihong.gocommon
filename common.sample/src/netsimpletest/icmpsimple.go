//icmpsimple.go
package netsimpletest

import (
	"bytes"
	"fmt"
	"net"
	"os"
)

func Simple1() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host")
		os.Exit(1)
	}
	service := os.Args[1]
	conn, err := net.Dial("ip4:icmp", service)
	checkError(err)
	var msg [512]byte
	msg[0] = 8 //echo
	msg[1] = 0 //code 0
	msg[2] = 0 //checksum
	msg[3] = 0 //checksum
	msg[4] = 0 //identi
}
