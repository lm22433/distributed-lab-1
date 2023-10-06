package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

type Message struct {
	sender  int
	message string
}

func acceptConns(ln net.Listener, conns chan net.Conn) {
	// Continuously accept a network connection from the Listener
	// and add it to the channel for handling connections.
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
		}
		conns <- conn
	}
}

func handleClient(client net.Conn, clientid int, msgs chan Message) {
	scanner := bufio.NewScanner(client)
	for scanner.Scan() {
		message := scanner.Text()
		msgs <- Message{sender: clientid, message: message}
	}
}

func main() {
	// Read in the network port we should listen on, from the commandline argument.
	// Default to port 8030
	portPtr := flag.String("port", ":8030", "port to listen on")
	flag.Parse()

	// Create a listener for TCP connections on the given port
	ln, err := net.Listen("tcp", *portPtr)
	if err != nil {
		fmt.Println("Error creating listener:", err)
		return
	}

	//Create a channel for connections
	conns := make(chan net.Conn)
	//Create a channel for messages
	msgs := make(chan Message)
	//Create a mapping of IDs to connections
	clients := make(map[int]net.Conn)

	//Start accepting connections
	go acceptConns(ln, conns)

	clientId := 0
	for {
		select {
		case conn := <-conns:
			// Assign a client ID
			clients[clientId] = conn
			// Start handling messages from this client
			go handleClient(conn, clientId, msgs)
			// Increment client id
			clientId++
		case msg := <-msgs:
			for id, conn := range clients {
				// Ignore the sender
				if id != msg.sender {
					_, err = fmt.Fprintln(conn, fmt.Sprintf("Client %d: %s", msg.sender, msg.message))
					if err != nil {
						fmt.Println("Error sending message:", err)
					}
				}
			}
		}
	}
}
