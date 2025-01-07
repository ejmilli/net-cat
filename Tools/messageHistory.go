package Tools

import "sync"

const maxHistory = 100

var (
	messageHistory []string
	historyLock    sync.Mutex
)

func AddMessageToHistory(msg string) {
	historyLock.Lock()
	defer historyLock.Unlock()

	if len(messageHistory) >= maxHistory {
		messageHistory = messageHistory[1:]
	}
	messageHistory = append(messageHistory, msg)
}

func GetMessageHistory() []string {
	historyLock.Lock()
	defer historyLock.Unlock()

	historyCopy := make([]string, len(messageHistory))
	copy(historyCopy, messageHistory)
	return historyCopy
}
