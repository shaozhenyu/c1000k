package main

import (
	"bufio"
	"fmt"
	//"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"time"
)

var quitSemaphore chan bool

func main() {

	count := 10
	if len(os.Args) > 1 {
		count, _ = strconv.Atoi(os.Args[1])
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	for i := 0; i < count; i++ {
		go Conn()
	}

	s := <-c
	fmt.Println("Got signal:", s)
}

func Conn() {
	conn, err := net.Dial("tcp", "192.168.100.156:8080")
	if err != nil {
		fmt.Println("conn err: ", err)
		return	
	}
	defer conn.Close()

	go messageRecived(conn)

	b := []byte("time\n")
	conn.Write(b)

	<-quitSemaphore
	fmt.Println("close ?????")
}

func messageRecived(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			quitSemaphore <- true
			break
		}
		time.Sleep(time.Second)
		b := []byte(msg)
		conn.Write(b)
	}
}
