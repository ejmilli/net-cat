package Tools

import "sync"

var (
	activeClients int
	clientLock    sync.Mutex
)

func GetActiveClients() int {
	clientLock.Lock()
	defer clientLock.Unlock()
	return activeClients
}

func AddActiveClients() {
	clientLock.Lock()
	activeClients++
	clientLock.Unlock()
}

func RemoveActiveClients() {
	clientLock.Lock()
	activeClients--
	clientLock.Unlock()
}
