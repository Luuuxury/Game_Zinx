package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// 1. 连接拨打服务器
	fmt.Println("Client start conn...")
	time.Sleep(1 * time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("Client Dial err, exit")
		return
	}
	// 2. 写数据
	for {
		_, err := conn.Write([]byte("Client Write Test !"))
		if err != nil {
			fmt.Println("Client conn Write error:", err)
			return
		}
		// 3. 服务器返回客户端的请求
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Server Call back (conn.Read) err:", err)
			return
		}
		fmt.Printf("Server call back: %s, cnt= %d \n", buf, cnt)
		time.Sleep(1 * time.Second)
	}
}
