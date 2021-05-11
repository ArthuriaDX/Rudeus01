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
	TestSudoCmd = "test" // for wotoTest plugin
	SudoSudoCmd = "sudo" // for wotoSudo plugin
	PatSudoCmd  = "pat"  // for wotoPat plugin
)

var sudoCMDList map[string]func(*tg.Message, pTools.Arg)

func sudoListInit() {
	if sudoCMDList != nil {
		return
	}

	sudoCMDList = make(map[string]func(*tg.Message, pTools.Arg))

	addTestSudoCMD()
	addSudoSudoCMD()
	addPatSudoCMD()
}

func addTestSudoCMD() {
	if sudoCMDList != nil {
		if sudoCMDList[TestSudoCmd] == nil {
			sudoCMDList[TestSudoCmd] = testCommandHandler
		}
	}
}

func addSudoSudoCMD() {
	if sudoCMDList != nil {
		if sudoCMDList[SudoSudoCmd] == nil {
			sudoCMDList[SudoSudoCmd] = sudoCommandHandler
		}
	}
}

func addPatSudoCMD() {
	if sudoCMDList != nil {
		if sudoCMDList[PatSudoCmd] == nil {
			sudoCMDList[PatSudoCmd] = patCommandHandler
		}
	}
}
