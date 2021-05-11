// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoActions

import (
	"github.com/ALiwoto/rudeus01/wotoPacks/interfaces"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/messages/textMessage"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var updatesMap map[int]bool

// shouldHangle will check if you should handle the
// update or not.
func shouldHandle(update *tg.Update) bool {
	if updatesMap == nil {
		updatesMap = make(map[int]bool)
	}

	if !updatesMap[update.UpdateID] {
		updatesMap[update.UpdateID] = true
	} else {
		return false
	}

	return true
}

// HandleMessage will handle the update comming from the telegram servers.
func HandleMessage(update *tg.Update, settings interfaces.WSettings) {
	if update.CallbackQuery != nil {
		settings.SendSudo(update.CallbackQuery.Data)
	}
	switch getMessageType(update) {
	case NONE:
		return
	case TEXT_MESSAGE:
		textMessage.HandleTextMessage(update.Message)
	default:
		return
	}
}
