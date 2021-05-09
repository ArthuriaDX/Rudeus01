// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package botCommands

import (
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/pTools"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/wotoMorse"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/wotoPat"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func toMorse_handler(message *tgbotapi.Message, args pTools.Arg) {
	wotoMorse.ToMorse_handler(message, args)
}

func fromMorse_handler(message *tgbotapi.Message, args pTools.Arg) {
	wotoMorse.FromMorse_handler(message, args)
}

func pat_handler(message *tgbotapi.Message, args pTools.Arg) {
	wotoPat.Pat_Handler(message, args)
}
