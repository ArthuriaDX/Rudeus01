// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoMorse

import (
	"log"
	"strings"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/pTools"
	ws "github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ToMorse_handler(message *tg.Message, args pTools.Arg) {
	args[wv.BaseIndex] = ws.EMPTY
	is_bin := args.HasFlag(BIN_FLAG)
	send_pv := args.HasFlag(PV_FLAG, PRIVATE_FLAG)
	is_reply := message.ReplyToMessage != nil
	var full, trl string
	if is_reply {
		if !ws.IsEmpty(&message.ReplyToMessage.Text) {
			full = message.ReplyToMessage.Text
		} else {
			return
		}
	} else {
		full = strings.Join(args, wv.SPACE_VALUE)
	}

	if is_bin {
		trl = TranslateToBinary(full)
	} else {
		trl = TranslateToMorse(full)
	}
	var str string
	if !ws.IsEmpty(&trl) {
		if is_bin {
			str = TO_BINARY_M + trl
		} else {
			str = TO_MORSE_M + trl
		}
	}
	sendMorse(message, &str, is_reply, send_pv)
}

func FromMorse_handler(message *tg.Message, args pTools.Arg) {
	args[wv.BaseIndex] = ws.EMPTY
	is_reply := message.ReplyToMessage != nil
	send_pv := args.HasFlag(PV_FLAG, PRIVATE_FLAG)
	var full string
	if is_reply {
		if !ws.IsEmpty(&message.ReplyToMessage.Text) {
			full = message.ReplyToMessage.Text
		} else {
			return
		}
	} else {
		full = strings.Join(args, wv.SPACE_VALUE)
	}
	trl := TranslateFromMorse(full)
	var str string
	if !ws.IsEmpty(&trl) {
		str = FROM_MORSE_M + trl
	}
	sendMorse(message, &str, is_reply, send_pv)
}

func sendMorse(message *tg.Message, morse *string, reply, pv bool) {
	settings := appSettings.GetExisting()
	api := settings.GetAPI()
	if settings == nil {
		return
	}
	var msg tg.MessageConfig
	if pv {
		msg = tg.NewMessage(message.From.ID, *morse)
	} else {
		msg = tg.NewMessage(message.Chat.ID, *morse)
	}
	if reply {
		r := message.ReplyToMessage
		if r != nil {
			msg.ReplyToMessageID = r.MessageID
		} else {
			msg.ReplyToMessageID = message.MessageID
		}
	} else {
		msg.ReplyToMessageID = message.MessageID
	}
	//msg.ParseMode = tgbotapi.ModeMarkdownV2
	if _, err := api.Send(msg); err != nil {
		log.Println(err)
		return
	}
}
