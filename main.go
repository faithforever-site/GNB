package main

import "net/http"

func main() {
	go startHTTPServer()      // HTTP 文件服务
	go startWebSocketServer() // WebSocket 聊天
	go startTCPServer()       // TCP 消息通信

	// 提供前端页面
	http.Handle("/", http.FileServer(http.Dir("static")))

	// 阻塞主线程
	select {}
}
