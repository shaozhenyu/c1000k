package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
)

var ch chan int = make(chan int)

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	for i := 0; i < 10; i++ {
		go Conn()
	}

	s := <-c
	fmt.Println("Got signal:", s)
}

func Conn() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	b := make([]byte, 10)
	n, err := conn.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b[:n]))
}

