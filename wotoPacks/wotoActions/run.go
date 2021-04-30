// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoActions

import (
	"net/http"
	"os"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func RunBot(_token string, c *gin.Context) {
	settings := appSettings.GetExisting()
	bot, err := tgbotapi.NewBotAPI(_token)
	settings.SetAPI(bot)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}

	bot.Debug = false
	for {
		_runOnce(bot, settings)
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
			_runRaw()
			continue
		}
		_runRaw()
		HandleMessage(&update, _settings)
		_runRaw()
	}
	RunNew()
}

func RunNew() {
	appSettings.App_exit()
	_runRaw()
}

func _runRaw() {
	go func() {
		url := os.Getenv(wotoValues.RUDEUS_URL_KEY)
		if !wotoSecurity.IsEmpty(&url) {
			_, _ = http.Get(url)
		}
	}()
}
