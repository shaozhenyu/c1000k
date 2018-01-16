package main

import (
	"log"
	"net"
	"time"
)

var send []byte = []byte("aaaaa")

func main() {
	ln, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	for {
		_, err := conn.Write(send)
		if err != nil {
			return
		}
		time.Sleep(10 * time.Second)
	}
}
