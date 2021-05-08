// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package textMessage

import (
	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/messages/textMessage/botCommands"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleTextMessage(message *tg.Message) {
	if message == nil {
		return
	}

	if botCommands.IsSudoCommand(&message.Text) {
		if appSettings.GetExisting().IsSudo(message.From.ID) {
			botCommands.HandleSudoCommand(message)
		}
		return
	} else if botCommands.IsCommand(&message.Text) {
		//log.Println("Is HERE!")
		botCommands.HandleCommand(message)
	} else {
		//log.Println("In ELSeE!")
	}

}
