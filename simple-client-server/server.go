package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	msg, _ := reader.ReadString('\n')
	fmt.Println(msg)
	fmt.Fprintln(conn, "OK")
}

func main() {
	ln, _ := net.Listen("tcp", "127.0.0.1:8030")
	for {
		conn, _ := ln.Accept()
		go handleConnection(conn)
	}
}
