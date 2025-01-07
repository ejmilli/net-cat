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

var (
	clients   = make(map[net.Conn]string)
	broadcast = make(chan string)
	mu        sync.Mutex
)

func HandleBroadcasts() {
	for {

		msg := <-broadcast
		AddMessageToHistory(msg)

		mu.Lock()
		for conn := range clients {
			fmt.Fprintln(conn, msg)
		}
		mu.Unlock()
	}
}

func HandleClient(conn net.Conn) {
	defer func() {
		RemoveActiveClients()
		mu.Lock()
		delete(clients, conn)
		mu.Unlock()
		conn.Close()
		broadcast <- "A user has left the chat."
	}()

	AddActiveClients()

	Penguin(conn)
	fmt.Fprint(conn, "[ENTER YOUR NAME]:")
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
