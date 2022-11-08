package main

import (
	"log"
	"net"
	"time"
)

func main() {
	// listen on 8080 port
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("service failed to start : ", err)

	}
	defer listen.Close()
	log.Println("service started successfully")
	// block waiting for user to connect
	accept, err := listen.Accept()
	if err != nil {
		log.Fatal(err)
	}
	// set timeout
	err = accept.SetDeadline(time.Now().Add(time.Second))
	if err != nil {
		log.Fatal(err)
	}
	// set read timeout
	err = accept.SetReadDeadline(time.Now().Add(time.Second))
	if err != nil {
		log.Fatal(err)
	}
	// set write timeout
	err = accept.SetWriteDeadline(time.Now().Add(time.Second))
	if err != nil {
		log.Fatal(err)
	}
}
