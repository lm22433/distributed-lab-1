package main

import (
	"fmt"
	"net"
)

func main() {
	msg := "Hello from Server!"
	conn, _ := net.Dial("tcp", "127.0.0.1:8030")
	fmt.Fprintln(conn, msg)
}
