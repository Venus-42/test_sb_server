package main

import (
	"log"
	"net"
)

func main() {
	// 尝试连接百度服务器
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	log.Println("Connect success!")
}
