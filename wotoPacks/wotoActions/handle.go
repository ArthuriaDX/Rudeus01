// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoActions

import (
	"log"
	"os"
	"strconv"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/messages/textMessage"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// shouldHangle will check if you should handle the
// update or not.
func shouldHangle(_update *tgbotapi.Update) bool {
	_strIndex := os.Getenv(wotoValues.INDEX_KEY)
	_strTotal := os.Getenv(wotoValues.TOTAL_INDEX_KEY)
	// get the index of this bot.
	// if the index is 1, then the bot will answer any
	// incoming update request
	_index, _err1 := strconv.Atoi(_strIndex)
	if _err1 != nil {
		log.Println(_err1)
		return false
	}
	_total, _err2 := strconv.Atoi(_strTotal)
	if _err2 != nil {
		log.Println(_err2)
		return false
	}
	_i := _update.UpdateID % _total
	return _i == _index
}

func HandleMessage(update *tgbotapi.Update, _settings *appSettings.AppSettings) {
	switch getMessageType(update) {
	case NONE:
		log.Println(*update)
	case TEXT_MESSAGE:
		textMessage.HandleTextMessage(update.Message, _settings)
	default:

	}
}
