package main

import (
	"log"
	"net/http"
)

// main 函数启动 HTTP/WebSocket 服务和 TCP 服务
func main() {
	// 创建 HTTP 路由
	mux := http.NewServeMux()

	// 文件上传/下载路由
	mux.HandleFunc("/upload", uploadHandler)
	mux.HandleFunc("/download/", downloadHandler)
	mux.HandleFunc("/tcp", tcpForwardHandler)
	// WebSocket 路由
	mux.HandleFunc("/ws", wsHandler)

	// 前端页面路由
	mux.Handle("/", http.FileServer(http.Dir("static")))

	// 启动 TCP 服务（独立端口）
	go startTCPServer()

	// 启动 HTTP + WebSocket 服务（8080端口）
	log.Println("HTTP + WebSocket server running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
