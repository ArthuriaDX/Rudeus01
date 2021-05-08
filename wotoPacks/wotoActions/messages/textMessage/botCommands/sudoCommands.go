// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package botCommands

import (
	"strings"

	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func IsSudoCommand(_text *string) bool {
	return strings.HasPrefix(*_text, wotoValues.SUDO_PREFIX1)
}

func HandleSudoCommand(message *tgbotapi.Message) {
	sudoListInit()
	text := strings.ToLower(message.Text)
	text = strings.TrimPrefix(text, wotoValues.SUDO_PREFIX1)
	text = strings.TrimPrefix(text, wotoValues.SPACE_VALUE)
	texts := strings.Split(text, wotoValues.SPACE_VALUE)

	event := sudoCMDList[texts[wotoValues.BaseIndex]]
	if event != nil {
		event(message, texts)
	}
}
