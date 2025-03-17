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
		mu.Lock()
		name := clients[conn]
		delete(clients, conn)
		mu.Unlock()

		conn.Close()
		broadcast <- fmt.Sprintf("%s has left the chat.", name)
		RemoveActiveClients()
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

	for strings.TrimSpace(name) == "" {
		fmt.Fprint(conn, "[ENTER YOUR NAME]:")
		reader := bufio.NewReader(conn)
		name, err = reader.ReadString('\n')

		if err != nil {
			log.Printf("Error reading name: %v", err)
			return
		}
	}

	name = strings.TrimSpace(name)

	mu.Lock()
	clients[conn] = name
	mu.Unlock()

	history := GetMessageHistory()
	for _, msg := range history {
		fmt.Fprintln(conn, msg)
	}

	broadcast <- fmt.Sprintf("%s has joined the chat.", name)

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		if msg == "\n" {
			continue
		}
		msg = strings.TrimSpace(msg)

		timestamp := time.Now().Format("2006-01-02 15:04:05")
		broadcast <- fmt.Sprintf("[%s][%s]: %s", timestamp, name, msg)
	}
}
