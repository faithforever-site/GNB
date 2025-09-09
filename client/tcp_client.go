package main

import (
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:9090")
	defer conn.Close()

	conn.Write([]byte("Hello TCP Server"))
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	fmt.Println("Server response:", string(buf[:n]))
}
