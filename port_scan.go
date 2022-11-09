package main

import (
	"flag"
	"log"
	"net"
	"time"
)

func scan(ip, port string) {
	conn, err := net.DialTimeout("tcp", ip+":"+port, time.Second*2)
	if err != nil {
		log.Fatalf("%s port not open", port)
	}
	defer conn.Close()
	log.Printf("%s port is open", port)
}
func main() {
	ip := flag.String("h", "127.0.0.1", "指定主机IP")
	port := flag.String("p", "9050", "指定端口")
	flag.Parse()
	scan(*ip, *port)
}
