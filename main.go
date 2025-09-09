package main

func main() {
	// 启动 HTTP 文件服务
	go startHTTPServer()

	// 启动 WebSocket 聊天服务
	go startWebSocketServer()

	// 启动 TCP 服务器
	go startTCPServer()

	// 阻塞主线程
	select {}
}
