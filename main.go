package main

import (
	"log"
	"net-cat/Tools"
)

func main() {

	ln := Tools.PortInput()
	// Goroutine to handle broadcasting
	go Tools.HandleBroadcasts()

	for {
		// Accept a new connection
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		// Handle the new client
		go Tools.HandleClient(conn)
	}
}
