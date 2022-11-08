package main

import (
	"log"
	"net"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	var buf = make([]byte, 1024)
	for {
		log.Println("start to read from conn")
		n, err := c.Read(buf)
		checkErr(err)
		log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
	}
}

func main() {
	// listen on 8080 port
	listen, err := net.Listen("tcp", ":8080")
	checkErr(err)
	defer listen.Close()
	log.Println("service started successfully!")
	for {
		conn, err := listen.Accept()
		checkErr(err)
		go handleConn(conn)
	}
}
