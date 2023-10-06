package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8030")
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		return
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter a message (or 'exit' to quit): ")
		scanner.Scan()
		msg := scanner.Text()

		if msg == "exit" {
			fmt.Println("Exiting client.")
			return
		}

		fmt.Fprintln(conn, msg)

		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading server response:", err)
			return
		}

		fmt.Print("Server response: " + response)
	}
}
