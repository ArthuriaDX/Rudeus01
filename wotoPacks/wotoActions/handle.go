// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoActions

import (
	"github.com/ALiwoto/rudeus01/wotoPacks/interfaces"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/messages/textMessage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// shouldHangle will check if you should handle the
// update or not.
func shouldHangle(_update *tgbotapi.Update) bool {

	return true
}

func HandleMessage(update *tgbotapi.Update, _settings interfaces.WSettings) {
	if update.CallbackQuery != nil {
		_settings.SendSudo(update.CallbackQuery.Data)
	}
	switch getMessageType(update) {
	case NONE:
		return
	case TEXT_MESSAGE:
		//log.Println("in TEXT switch! "+update.Message.Text, _settings)
		textMessage.HandleTextMessage(update.Message)
	default:
		return
	}
}
