package main

import (
	"log"
	"net"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	// try to connect local 1234 port
	udpaddr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:8080")
	checkErr(err)
	// establish connect
	conn, err := net.DialUDP("udp", nil, udpaddr)
	checkErr(err)
	defer conn.Close()
	log.Println("successfully connected")
	// send message
	_, err = conn.Write([]byte("Hello, world!\r\n"))
	checkErr(err)
	// accept message
	var buf = make([]byte, 1024)
	n, err := conn.Read(buf)
	checkErr(err)
	log.Printf("accept information length:%d; info:%s", n, string(buf[:n]))
}
