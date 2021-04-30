//+build amd64, windows, linux

package wotoPacks

import (
	"fmt"
	"net"
	"time"
)

//goland:noinspection GoUnusedExportedFunction,GoSnakeCaseUsage
func ServerDaytime(_port string){
	service := "localhost:" + _port
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	CheckError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	CheckError(err)
	for {
		fmt.Println("came to loop!")
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("whoa! I got some clients here!")
		daytime := time.Now().String()
		_, _ = conn.Write([]byte(daytime)) // don't care about return value
		_ = conn.Close()
		// we're finished with this client
	}
}

