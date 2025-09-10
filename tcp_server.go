package main

import (
	"fmt"
	"net"
)

// 启动 TCP 服务（独立端口 9090）
func startTCPServer() {
	ln, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println("TCP listen error:", err)
		return
	}
	defer ln.Close()
	fmt.Println("TCP server running on :9090")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("TCP accept error:", err)
			continue
		}
		go handleTCPConn(conn) // 并发处理客户端
	}
}

// 处理单个 TCP 客户端连接
func handleTCPConn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("TCP read error:", err)
		return
	}
	fmt.Println("TCP Received:", string(buf[:n]))
	// 回应客户端
	conn.Write([]byte("TCP message received"))
}
