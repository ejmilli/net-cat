package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

var (
	clients   = make(map[net.Conn]string) // Map of active clients
	broadcast = make(chan string)         // Channel for broadcasting messages
	mu        sync.Mutex                  // Mutex to protect shared resources
)

func main() {
	// Start listening on port 2525
	ln, err := net.Listen("tcp", ":2525")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer ln.Close()
	log.Println("Chat server started on port 2525. Connect with ncat localhost 2525.")

	// Goroutine to handle broadcasting
	go handleBroadcasts()

	for {
		// Accept a new connection
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		// Handle the new client
		go handleClient(conn)
	}
}

func handleBroadcasts() {
	for {
		// Wait for a message to broadcast
		msg := <-broadcast

		// Send the message to all connected clients
		mu.Lock()
		for conn := range clients {
			fmt.Fprintln(conn, msg)
		}
		mu.Unlock()
	}
}

func handleClient(conn net.Conn) {
	defer func() {
		// Remove the client when they disconnect
		mu.Lock()
		delete(clients, conn)
		mu.Unlock()
		conn.Close()
		broadcast <- "A user has left the chat."
	}()


	
	// Ask the client for their name
	fmt.Fprint(conn, "Enter your name: ")
	reader := bufio.NewReader(conn)
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Error reading name: %v", err)
		return
	}
	name = strings.TrimSpace(name)

	// Add the client to the list
	mu.Lock()
	clients[conn] = name
	mu.Unlock()

	broadcast <- fmt.Sprintf("%s has joined the chat.", name)

	// Read messages from the client
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading message from %s: %v", name, err)
			return
		}
		msg = strings.TrimSpace(msg)

		// Add timestamp to the message
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		broadcast <- fmt.Sprintf("[%s] %s: %s", timestamp, name, msg)
	}
}