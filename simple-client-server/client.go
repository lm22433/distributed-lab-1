package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Prompt the user for a message
	fmt.Print("Enter the message: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	msg := scanner.Text()

	conn, err := net.Dial("tcp", "127.0.0.1:8030")
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		return
	}
	defer conn.Close()

	fmt.Fprintln(conn, msg)
}
