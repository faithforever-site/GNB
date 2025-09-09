**GNB
network application by golang
# 项目说明

- **HTTP 文件上传/下载**

- **WebSocket 聊天**

- **TCP 简单消息通信**


```
# 项目结构

gnb/
│
├─ main.go            # 启动所有服务
├─ http_server.go     # HTTP 文件上传下载
├─ websocket_server.go# WebSocket 聊天
├─ client                    # 单独用一个文件夹放客户端文件
     └─ tcp_client.go      # TCP 测试客户端
├─ tcp_server.go      #  TCP 消息通信服务端
└─ upload/            # 上传文件存放目录
```
