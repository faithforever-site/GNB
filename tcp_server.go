package main

import (
	"fmt"
	"net"
)

func startTCPServer() {
	ln, _ := net.Listen("tcp", ":9090")
	defer ln.Close()
	fmt.Println("TCP server running on :9090")

	for {
		conn, _ := ln.Accept()
		go handleTCPConn(conn)
	}
}

func handleTCPConn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	fmt.Println("TCP Received:", string(buf[:n]))
	conn.Write([]byte("TCP message received"))
}
