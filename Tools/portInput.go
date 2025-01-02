package Tools

import (
	"log"
	"os"
	"strconv"
)

func PortInput() string {
	var port string
	if len(os.Args) == 2 {
		port = os.Args[1]

		portINT, err := strconv.Atoi(port)
		if err != nil {
			log.Fatalf("This is not a port number: %v", err)
		}

		if portINT < 1024 || portINT > 65535 {
			log.Fatalf("The range of the port number: 1024-65535. This is not a valid port number: %v", err)
		}
	}

	if len(os.Args) == 1 {
		port = "8989" // Default port
	}

	return port
}
