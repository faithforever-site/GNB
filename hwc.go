package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	// 允许跨域连接，开发阶段用，生产中可以限制 Origin
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]bool)

// WebSocket 处理函数
func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil) // HTTP 升级到 WebSocket
	if err != nil {
		log.Println(err)
		return
	}
	clients[conn] = true
	defer func() {
		delete(clients, conn)
		conn.Close()
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break // 客户端关闭或出错
		}
		broadcast(msg) // 转发给所有客户端
	}
}

// WebSocket 广播消息给所有客户端
func broadcast(msg []byte) {
	for client := range clients {
		client.WriteMessage(websocket.TextMessage, msg)
	}
}

// HTTP 文件上传处理
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Write([]byte("请使用 POST 上传文件"))
		return
	}

	// 创建上传目录
	os.MkdirAll("upload", os.ModePerm)

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	dst, err := os.Create(filepath.Join("upload", header.Filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// 保存文件
	io.Copy(dst, file)
	w.Write([]byte("Upload successful"))
}

// HTTP 文件下载处理
func downloadHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Path[len("/download/"):]
	filepath := filepath.Join("upload", filename)
	http.ServeFile(w, r, filepath)
}
func tcpForwardHandler(w http.ResponseWriter, r *http.Request) {
	// 只允许 POST
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 读取浏览器消息
	msg, _ := io.ReadAll(r.Body)

	// 建立 TCP 连接到 TCP 服务
	conn, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		http.Error(w, "TCP server not reachable", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	// 发送消息到 TCP
	conn.Write(msg)

	// 读取 TCP 服务返回
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)

	// 返回给浏览器
	w.Write(buf[:n])
}

//浏览器 (只能发 HTTP/HTTPS/WebSocket)
//↓
//HTTP 服务 (Go)
//↓  (这里用 Go 代码做代理/转发)
//TCP 服务 (Go，监听 9090)
//↓
//HTTP 服务把 TCP 响应转回来
//↓
//浏览器收到 HTTP 响应，显示出来
