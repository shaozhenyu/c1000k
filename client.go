package main

import (
	"fmt"
	//"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"time"
)

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

	for {
		b := make([]byte, 10)
		_, err = conn.Read(b)
		if err != nil {
			fmt.Println("read err: ", err)
			return
		}
		time.Sleep(10 * time.Second)
	}
	fmt.Println("close ?????")
}
