package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 连接 TCP 服务器
	conn, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		fmt.Println("连接失败:", err)
		return
	}
	defer conn.Close()

	fmt.Println("已连接到 TCP 服务器，输入消息并回车 (输入 exit 退出)：")

	reader := bufio.NewReader(os.Stdin)
	for {
		// 从控制台读取输入
		fmt.Print("你: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Println("退出客户端")
			return
		}

		// 发送消息到 TCP 服务端
		_, err := conn.Write([]byte(text))
		if err != nil {
			fmt.Println("发送失败:", err)
			return
		}

		// 读取服务端回应
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("读取失败:", err)
			return
		}
		fmt.Println("服务器:", string(buf[:n]))
	}
}
