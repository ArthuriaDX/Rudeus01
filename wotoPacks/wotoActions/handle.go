// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoActions

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// shouldHangle will check if you should handle the
// update or not.
func shouldHangle(_update *tgbotapi.Update) bool {

	return false
}

func Handle(_message *tgbotapi.Message) {

}
