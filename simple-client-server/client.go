package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func read(conn net.Conn) {
	reader := bufio.NewReader(conn)
	msg, _ := reader.ReadString('\n')
	fmt.Printf(msg)
}

func main() {
	fmt.Print("Enter the message: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	msg := scanner.Text()

	conn, _ := net.Dial("tcp", "127.0.0.1:8030")
	fmt.Fprintln(conn, msg)
	read(conn)
}
