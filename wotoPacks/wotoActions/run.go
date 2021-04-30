// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoActions

import (
	"log"
	"net/http"
	"os"

	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/messages/textMessage"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func RunBot(_token string, c *gin.Context) {
	bot, err := tgbotapi.NewBotAPI(_token)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}

	bot.Debug = false
	// Create a new UpdateConfig struct with an offset of 0. Offsets are used
	// to make sure Telegram knows we've handled previous values and we don't
	// need them repeated.
	updateConfig := tgbotapi.NewUpdate(wotoValues.BaseIndex)

	// Tell Telegram we should wait up to 30 seconds on each request for an
	// update. This way we can get information just as quickly as making many
	// frequent requests without having to send nearly as many.
	updateConfig.Timeout = wotoValues.BaseTimeOut

	// Start polling Telegram for updates.
	updates := bot.GetUpdatesChan(updateConfig)

	// Let's go through each update that we're getting from Telegram.
	for update := range updates {

		// check if the current application is allowed to handle the update
		// request or not.
		if !shouldHangle(&update) {
			continue
		}

		_type := getMessageType(&update)

		switch _type {
		case NONE:
			log.Println(NONE)
			continue
		case TEXT_MESSAGE:
			textMessage.HandleTextMessage(update.Message)
		default:
			continue
		}

		// Telegram can send many types of updates depending on what your Bot
		// is up to. We only want to look at messages for now, so we can
		// discard any other updates.
		if update.Message == nil {

			continue
		}

		// Now that we know we've gotten a new message, we can construct a
		// reply! We'll take the Chat ID and Text from the incoming message
		// and use it to create a new message.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		// We'll also say that this message is a reply to the previous message.
		// For any other specifications than Chat ID or Text, you'll need to
		// set fields on the `MessageConfig`.
		msg.ReplyToMessageID = update.Message.MessageID

		// Okay, we're sending our message off! We don't care about the message
		// we just sent, so we'll discard it.
		if _, err := bot.Send(msg); err != nil {
			// Note that panics are a bad way to handle errors. Telegram can
			// have service outages or network errors, you should retry sending
			// messages or more gracefully handle failures.
			RunNew()
			return
		}
	}

}

func RunNew() {
	url := os.Getenv(wotoValues.RUDEUS_URL_KEY)
	if !wotoSecurity.IsEmpty(&url) {
		_, _ = http.Get(url)
	}
}
