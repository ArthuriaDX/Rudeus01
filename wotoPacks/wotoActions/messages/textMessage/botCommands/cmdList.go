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
	TMORSe_CMD = "tmorse"
	FMORSE_CMD = "fmorse"
)

var cmdList map[string]func(*tgbotapi.Message, []string)

func cmdListInit() {
	if cmdList != nil {
		return
	}
	log.Println("In Init!")
	cmdList = make(map[string]func(*tgbotapi.Message, []string))
	add_morseCmdList()
}

func add_morseCmdList() {
	if cmdList != nil {
		if cmdList[TMORSe_CMD] == nil {
			cmdList[TMORSe_CMD] = toMorse_handler
		}
		if cmdList[FMORSE_CMD] == nil {
			cmdList[FMORSE_CMD] = fromMorse_handler
		}
	}
}
