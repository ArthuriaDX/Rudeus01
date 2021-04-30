// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoActions

import (
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageType uint8

const (
	NONE                     = 0
	TEXT_MESSAGE MessageType = 1
	GIF_MESSAGE  MessageType = 2
)

func getMessageType(_update *tgbotapi.Update) MessageType {
	if _checkTextMessage(_update) {
		return TEXT_MESSAGE
	}

	if _checkGifMessage(_update) {
		return GIF_MESSAGE
	}

	return NONE
}

func _checkTextMessage(_update *tgbotapi.Update) bool {
	_msg := _update.Message
	if _msg.MessageID == wotoValues.BaseIndex {
		return false
	}
	if _msg.From == nil {
		return false
	}
	if _msg.Chat == nil {
		return false
	}
	if _msg.Date == wotoValues.BaseIndex {
		return false
	}
	if wotoSecurity.IsEmpty(&_msg.Text) {
		return false
	}
	if _msg.Animation != nil {
		return false
	}
	if _msg.Document != nil {
		return false
	}
	if !wotoSecurity.IsEmpty(&_msg.Caption) {
		return false
	}
	if _msg.Photo != nil {
		return false
	}
	if _msg.Video != nil {
		return false
	}
	if _msg.Game != nil {
		return false
	}

	return true
}

func _checkGifMessage(_update *tgbotapi.Update) bool {

	return false
}
