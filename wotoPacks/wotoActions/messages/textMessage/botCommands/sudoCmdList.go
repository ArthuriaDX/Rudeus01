// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

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
	add_testSudoCMD()
}

func add_testSudoCMD() {
	if sudoCMDList != nil {
		if sudoCMDList[TEST_SUDO_CMD] == nil {
			sudoCMDList[TEST_SUDO_CMD] = testCommandHandle
		}
	}
}
