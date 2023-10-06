package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from client:", err)
			return
		}
		fmt.Print("Received: " + msg)
		fmt.Fprintln(conn, "Server: Acknowledged - "+msg)
	}
}

func main() {
	ln, _ := net.Listen("tcp", "127.0.0.1:8030")
	defer ln.Close()
	fmt.Println("Server is listening on port 8030")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
