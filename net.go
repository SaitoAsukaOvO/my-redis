package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func ListenAndServe(address string) {
	// bind listen address with tcp
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(fmt.Sprintf("listen err: %v", err))
	}
	defer listener.Close()
	log.Println(fmt.Sprintf("bind: %s, start listening...", address))

	for {
		// Accept is blocked until a new connection is established or the listener is interrupted.
		conn, err := listener.Accept()
		if err != nil {
			// usually caused by closed listener.
			log.Fatal(fmt.Sprintf("accept err: %v", err))
		}
		// open a new goroutine to handle
		go Handle(conn)
	}
}

func Handle(conn net.Conn) {
	// use bufio to provide a buffer
	reader := bufio.NewReader(conn)
	for {
		// ReadString will be blocked until the separator '\n' comes in.
		msg, err := reader.ReadString('\n')
		if err != nil {
			// 通常遇到的错误是连接中断或被关闭，用io.EOF表示
			if err == io.EOF {
				log.Println("connection close")
			} else {
				log.Println(err)
			}
			return
		}
		b := []byte(msg)
		// send message to client
		conn.Write(b)
	}
}

func main() {
	ListenAndServe(":8000")
}
