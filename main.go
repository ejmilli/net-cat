package main

import (
	"fmt"
	"log"
	"net-cat/Tools"
)

func main() {

	ln := Tools.PortInput()
	const maxClients = 10

	defer ln.Close()
	
	// Goroutine to handle broadcasting
	go Tools.HandleBroadcasts()

	for {
		// Accept a new connection
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		if Tools.GetActiveClients() >= maxClients {
			fmt.Fprintln(conn, "Server is full. Try again later.")
			conn.Close()
			continue
		}

		// Handle the new client
		go Tools.HandleClient(conn)
	}
}
