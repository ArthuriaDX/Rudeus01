// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package botCommands

import (
	"log"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func testCommandHandle(message *tgbotapi.Message, args []string) {
	settings := appSettings.GetExisting()
	//log.Println("In EVENT!")
	if settings == nil {
		return
	}
	var str string
	if len(args) <= wotoValues.BaseOneIndex {
		str = "no test session provided!"
	} else {
		str = "couldn't found test session : " + args[wotoValues.BaseOneIndex]
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, str)
	msg.ReplyToMessageID = message.MessageID
	if _, err := settings.GetAPI().Send(msg); err != nil {
		log.Println(err)
		return
	}
}
