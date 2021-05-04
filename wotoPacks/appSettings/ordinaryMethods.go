// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package appSettings

import (
	"log"
	"net/http"

	ws "github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (_s *AppSettings) InvalidateAPI() bool {
	if _s._apiObtain != nil {
		_s.isGlobal = !ws.IsEmpty(&_s.tObt)
		b := _s._apiObtain(_s)
		if !b {
			if _s._nextChild != nil {
				_s._nextChild(_s.url, _s)
			}
		}
		return b
	} else {
		_s.isGlobal = false
		return _s.isGlobal
	}
}

func App_enter() {
	_alreadyRunning = true
}

func App_exit() {
	_alreadyRunning = false
}

func (_s *AppSettings) RunNext() {
	if !ws.IsEmpty(&_s.nextUrl) {
		go func() {
			_, _err := http.Get(_s.nextUrl)
			if _err != nil {
				log.Println(_err, _s.nextUrl)
			}
		}()
	}
}

// SendSudo, will send a message to the superuser of the bot.
func (_s *AppSettings) SendSudo(str string) {
	// Now that we know we've gotten a new message, we can construct a
	// reply! We'll take the Chat ID and Text from the incoming message
	// and use it to create a new message. 1341091260
	msg := tg.NewMessage(wv.SUDO, str)

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
