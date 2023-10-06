package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func read(conn net.Conn) {
	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from server:", err)
		}
		fmt.Print(msg)
	}
}

func write(conn net.Conn) {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		msg := scanner.Text()
		_, err := fmt.Fprintln(conn, msg)
		if err != nil {
			fmt.Println("Error sending message:", err)
		}
	}
}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()

	// Attempt to connect to the server and handle the error.
	conn, err := net.Dial("tcp", *addrPtr)
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
	}

	// Asynchronously start both the read and write goroutines.
	go read(conn)
	go write(conn)

	for {
	}
}
