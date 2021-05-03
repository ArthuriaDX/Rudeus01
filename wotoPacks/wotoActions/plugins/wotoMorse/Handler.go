// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoMorse

import (
	"log"
	"strings"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ToMorse_handler(message *tgbotapi.Message, args []string) {
	args[wotoValues.BaseIndex] = wotoSecurity.EMPTY
	is_reply := message.ReplyToMessage != nil
	var full string
	if is_reply {
		if !wotoSecurity.IsEmpty(&message.ReplyToMessage.Text) {
			full = message.ReplyToMessage.Text
		} else {
			return
		}
	} else {
		full = strings.Join(args, wotoValues.SPACE_VALUE)
	}
	trl := TranslateToMorse(full)
	var str string
	if !wotoSecurity.IsEmpty(&trl) {
		str = "Done! the morse code is: \n" + trl
	}
	if is_reply {
		sendMorse(message.ReplyToMessage, &trl)
	} else {
		sendMorse(message, &str)
	}
}

func FromMorse_handler(message *tgbotapi.Message, args []string) {
	args[wotoValues.BaseIndex] = wotoSecurity.EMPTY
	is_reply := message.ReplyToMessage != nil
	var full string
	if is_reply {
		if !wotoSecurity.IsEmpty(&message.ReplyToMessage.Text) {
			full = message.ReplyToMessage.Text
		} else {
			return
		}
	} else {
		full = strings.Join(args, wotoValues.SPACE_VALUE)
	}
	trl := TranslateFromMorse(full)
	var str string
	if !wotoSecurity.IsEmpty(&trl) {
		str = "Done! the morse code is: \n" + trl
	}
	if is_reply {
		sendMorse(message.ReplyToMessage, &trl)
	} else {
		sendMorse(message, &str)
	}
}

func sendMorse(message *tgbotapi.Message, morse *string) {
	settings := appSettings.GetExisting()
	if settings == nil {
		return
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, *morse)
	msg.ReplyToMessageID = message.MessageID
	//msg.ParseMode = tgbotapi.ModeMarkdownV2
	if _, err := settings.GetAPI().Send(msg); err != nil {
		log.Println(err)
		return
	}
}
