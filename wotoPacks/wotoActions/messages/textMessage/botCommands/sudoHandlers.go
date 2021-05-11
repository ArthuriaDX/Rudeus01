// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package botCommands

import (
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/pTools"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/wotoPat"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/wotoSudo"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/wotoTest"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func testCommandHandler(message *tg.Message, args pTools.Arg) {
	wotoTest.TestCommandHandle(message, args)
}

func sudoCommandHandler(message *tg.Message, args pTools.Arg) {
	wotoSudo.Sudo_handler(message, args)
}

func patCommandHandler(message *tg.Message, args pTools.Arg) {
	wotoPat.PatSHandler(message, args)
}
