// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package botCommands

import (
	"log"
	"strings"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func IsSudoCommand(_text *string) bool {
	return strings.HasPrefix(*_text, wotoValues.SUDO_PREFEX1)
}

func HandleSudoCommand(message *tgbotapi.Message) {
	sudoListInit()
	text := strings.ToLower(message.Text)
	text = strings.TrimPrefix(text, wotoValues.SUDO_PREFEX1)
	text = strings.TrimPrefix(text, wotoValues.SPACE_VALUE)
	texts := strings.Split(text, wotoValues.SPACE_VALUE)
	//log.Println("before event : ", texts)
	event := sudoCMDList[texts[wotoValues.BaseIndex]]
	if event != nil {
		event(message, texts)
	}
}

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
	// Now that we know we've gotten a new message, we can construct a
	// reply! We'll take the Chat ID and Text from the incoming message
	// and use it to create a new message. 1341091260
	msg := tgbotapi.NewMessage(message.Chat.ID, str)

	//msg.ParseMode = tgbotapi.ModeMarkdownV2

	// We'll also say that this message is a reply to the previous message.
	// For any other specifications than Chat ID or Text, you'll need to
	// set fields on the `MessageConfig`.
	msg.ReplyToMessageID = message.MessageID
	if _, err := settings.GetAPI().Send(msg); err != nil {
		// Note that panics are a bad way to handle errors. Telegram can
		// have service outages or network errors, you should retry sending
		// messages or more gracefully handle failures.
		log.Println(err)
		return
	}
}
