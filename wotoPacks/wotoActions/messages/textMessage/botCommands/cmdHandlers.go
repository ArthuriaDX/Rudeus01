// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package botCommands

import (
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/wotoMorse"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func toMorse_handler(message *tgbotapi.Message, args []string) {
	wotoMorse.ToMorse_handler(message, args)
}

func fromMorse_handler(message *tgbotapi.Message, args []string) {
	wotoMorse.FromMorse_handler(message, args)
}
