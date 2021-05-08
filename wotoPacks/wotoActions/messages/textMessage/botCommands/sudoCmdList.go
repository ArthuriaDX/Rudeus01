// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package botCommands

import (
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/pTools"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// sudo commands
const (
	TEST_SUDO_CMD = "test"
	SUDO_SUDO_CMD = "sudo"
)

var sudoCMDList map[string]func(*tg.Message, pTools.Arg)

func sudoListInit() {
	if sudoCMDList != nil {
		return
	}
	//log.Println("In Init!")
	sudoCMDList = make(map[string]func(*tg.Message, pTools.Arg))
	add_testSudoCMD()
	add_sudoSudoCMD()
}

func add_testSudoCMD() {
	if sudoCMDList != nil {
		if sudoCMDList[TEST_SUDO_CMD] == nil {
			sudoCMDList[TEST_SUDO_CMD] = testCommandHandle
		}
	}
}

func add_sudoSudoCMD() {
	if sudoCMDList != nil {
		if sudoCMDList[SUDO_SUDO_CMD] == nil {
			sudoCMDList[SUDO_SUDO_CMD] = sudoCommandHandle
		}
	}
}
