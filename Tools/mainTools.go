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
	clients   = make(map[net.Conn]string) // Map of active clients
	broadcast = make(chan string)         // Channel for broadcasting messages
	mu        sync.Mutex                  // Mutex to protect shared resources
)

func HandleBroadcasts() {
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

func HandleClient(conn net.Conn) {
	defer func() {
		// Remove the client when they disconnect
		mu.Lock()
		delete(clients, conn)
		mu.Unlock()
		conn.Close()
		broadcast <- "A user has left the chat."
	}()

	fmt.Fprintln(conn, "Welcome to TCP-Chat!")
	fmt.Fprintln(conn, "         _nnnn_")
	fmt.Fprintln(conn, "        dGGGGMMb")
	fmt.Fprintln(conn, "       @p~qp~~qMb")
	fmt.Fprintln(conn, "       M|@||@) M|")
	fmt.Fprintln(conn, "       @,----.JM|")
	fmt.Fprintln(conn, "      JS^\\__/  qKL")
	fmt.Fprintln(conn, "     dZP        qKRb")
	fmt.Fprintln(conn, "    dZP          qKKb")
	fmt.Fprintln(conn, "   fZP            SMMb")
	fmt.Fprintln(conn, "   HZM            MMMM")
	fmt.Fprintln(conn, "   FqM            MMMM")
	fmt.Fprintln(conn, " __| \".        |\\dS\"qML")
	fmt.Fprintln(conn, " |    `.       | `' \\Zq")
	fmt.Fprintln(conn, "_)      \\.___.,|     .'")
	fmt.Fprintln(conn, "\\____   )MMMMMP|   .'")
	fmt.Fprintln(conn, "     `-'       `--'")
	fmt.Fprintln(conn)

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
