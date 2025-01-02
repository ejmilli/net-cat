package main

import (
	"log"
	"net"
	"net-cat/Tools"
)

func main() {

	port := Tools.PortInput()

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer ln.Close()

	log.Println("Chat server started on port", port, ". Connect with ncat localhost", port, ".")

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
