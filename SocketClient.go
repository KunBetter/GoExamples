// SocketClient
package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
)

func main() {
	remoteHost := "localhost"
	remotePort := 9999
	tcpAddr, err := net.ResolveTCPAddr("tcp4", remoteHost+":"+strconv.Itoa(remotePort))
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	_, err := conn.Write([]byte("Head / HTTP/1.0\r\n\r\n"))
	checkError(err)
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}
