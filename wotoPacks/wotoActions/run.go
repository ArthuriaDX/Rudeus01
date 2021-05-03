// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoActions

import (
	"log"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoDB"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func RunBot(_settings *appSettings.AppSettings, _client *wotoDB.WotoClient) {
	bot, err := tgbotapi.NewBotAPI(_settings.GetObt())
	_settings.SetAPI(bot)
	if err != nil {
		log.Fatal(err)
		return
	}

	bot.Debug = false
	_client.PingClientDB(true)

	for {
		_runOnce(bot, _settings)
	}
}

func _runOnce(_bot *tgbotapi.BotAPI, _settings *appSettings.AppSettings) {
	// Create a new UpdateConfig struct with an offset of 0. Offsets are used
	// to make sure Telegram knows we've handled previous values and we don't
	// need them repeated.
	updateConfig := tgbotapi.NewUpdate(wotoValues.BaseIndex)

	// Tell Telegram we should wait up to 30 seconds on each request for an
	// update. This way we can get information just as quickly as making many
	// frequent requests without having to send nearly as many.
	updateConfig.Timeout = wotoValues.BaseTimeOut

	// Start polling Telegram for updates.
	updates := _bot.GetUpdatesChan(updateConfig)

	// Let's go through each update that we're getting from Telegram.
	for update := range updates {
		// check if the current application is allowed to handle the update
		// request or not.
		if !shouldHangle(&update) {
			continue
		}
		go HandleMessage(&update, _settings)
	}
}

// SendSudo, will send a message to the superuser of the bot.
func SendSudo(str string, _settings *appSettings.AppSettings) {
	// Now that we know we've gotten a new message, we can construct a
	// reply! We'll take the Chat ID and Text from the incoming message
	// and use it to create a new message. 1341091260
	msg := tgbotapi.NewMessage(1341091260, str)

	// We'll also say that this message is a reply to the previous message.
	// For any other specifications than Chat ID or Text, you'll need to
	// set fields on the `MessageConfig`.
	if _, err := _settings.GetAPI().Send(msg); err != nil {
		// Note that panics are a bad way to handle errors. Telegram can
		// have service outages or network errors, you should retry sending
		// messages or more gracefully handle failures.
		log.Println(err)
		return
	}
}
