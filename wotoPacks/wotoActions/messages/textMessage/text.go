// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package textMessage

import (
	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleTextMessage(_message *tgbotapi.Message, _settings *appSettings.AppSettings) {
	if _message == nil {
		return
	}
	_str := "_Oh Yeah_" +
		"I've got a new `Message!` this one: \n" +
		_message.Text + "\n" +
		"from : @" +
		_message.From.UserName + "ArthuriaDX!"
	// Now that we know we've gotten a new message, we can construct a
	// reply! We'll take the Chat ID and Text from the incoming message
	// and use it to create a new message.
	msg := tgbotapi.NewMessage(_message.Chat.ID, _str)

	// We'll also say that this message is a reply to the previous message.
	// For any other specifications than Chat ID or Text, you'll need to
	// set fields on the `MessageConfig`.
	msg.ReplyToMessageID = _message.MessageID
	if _, err := _settings.GetAPI().Send(msg); err != nil {
		// Note that panics are a bad way to handle errors. Telegram can
		// have service outages or network errors, you should retry sending
		// messages or more gracefully handle failures.
		_settings.RunNew()
		return
	}
}
