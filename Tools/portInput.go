package Tools

import (
	"log"
	"net"
	"os"
	"strconv"
)

func PortInput() net.Listener {
	var port string
	if len(os.Args) == 2 {
		port = os.Args[1]

		portINT, err := strconv.Atoi(port)
		if err != nil {
			log.Fatalf("This is not a port number: %v", err)
		}

		if portINT < 1024 || portINT > 65535 {
			log.Fatalf("The range of the port number: 1024-65535. This is not a valid port number: %v", portINT )
		}
	}

	if len(os.Args) == 1 {
		port = "8989" // Default port
	}

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	log.Println("Chat server started on port", port, ". Connect with nc localhost", port, ".")

	return ln
}
