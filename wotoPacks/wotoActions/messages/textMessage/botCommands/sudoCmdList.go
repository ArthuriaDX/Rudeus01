package botCommands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// sudo commands
const (
	TEST_SUDO_CMD = "test"
)

var sudoCMDList map[string]func(*tgbotapi.Message, []string)

func sudoListInit() {
	if sudoCMDList != nil {
		return
	}
	log.Println("In Init!")
	sudoCMDList = make(map[string]func(*tgbotapi.Message, []string))
	if sudoCMDList[TEST_SUDO_CMD] == nil {
		sudoCMDList[TEST_SUDO_CMD] = testCommandHandle
	}
}
