package Tools

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

// Global Variables
var (
	clients   = make(map[net.Conn]string)
	broadcast = make(chan string)
	mu        sync.Mutex
)

// HandleClient: Manages individual client connection
func HandleClient(conn net.Conn) {

	AddActiveClients()
	Penguin(conn) // Assuming Penguin() is some kind of welcome message

	// Prompt for user name
	fmt.Fprint(conn, "[ENTER YOUR NAME]:")
	reader := bufio.NewReader(conn)

	name, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Error reading name: %v", err)
		return
	}

	fmt.Println("conn:", conn)
	fmt.Println("reader", reader)
	fmt.Println("name:", name)
	// Ensure a valid name is entered
	for name == "\n" {
		fmt.Fprintln(conn, "type your name bitch:")
		name, err = reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading name: %v", err)
			return
		}
	}
	name = strings.TrimSpace(name)

	// Add client to active list
	mu.Lock()
	clients[conn] = name
	mu.Unlock()

	// Send chat history to the new client
	history := GetMessageHistory()
	for _, msg := range history {
		fmt.Fprintln(conn, msg)
	}

	// Announce the new user to others
	broadcast <- fmt.Sprintf("%s has joined the chat.", name)

	defer func() {
		mu.Lock()
		name := clients[conn]
		delete(clients, conn)
		mu.Unlock()
		conn.Close()
		broadcast <- fmt.Sprintf("%s has left the chat.", name)
		RemoveActiveClients()
	}()

	// Process incoming messages from the client
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		if msg == "\n" {
			continue
		}
		msg = strings.TrimSpace(msg)

		// Format and send the message to the broadcast channel
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		broadcast <- fmt.Sprintf("[%s][%s]: %s", timestamp, name, msg)
	}
	
}

// HandleBroadcasts: Processes and sends messages to all clients
func HandleBroadcasts() {
	for {
		msg := <-broadcast

		// Add the message to the chat history
		AddMessageToHistory(msg)

		// Send the message to all connected clients
		mu.Lock()
		for conn := range clients {
			fmt.Fprintln(conn, msg)
		}
		mu.Unlock()
	}
}


