// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoActions

import (
	ws "github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageType uint8

const (
	NONE                       = 0
	CALLBACK_QUERY MessageType = 1
	TEXT_MESSAGE   MessageType = 2
	GIF_MESSAGE    MessageType = 3
	PHOTO_MESSAGE  MessageType = 4
)

func getMessageType(_update *tg.Update) MessageType {

	if _checkCallBackQuery(_update) {
		return CALLBACK_QUERY
	}

	if _checkTextMessage(_update) {
		return TEXT_MESSAGE
	}

	if _checkGifMessage(_update) {
		return GIF_MESSAGE
	}

	if _checkPhotoMessage(_update) {
		return PHOTO_MESSAGE
	}

	return NONE
}

func _checkCallBackQuery(_update *tg.Update) bool {
	return _update.CallbackQuery != nil
}

func _checkTextMessage(_update *tg.Update) bool {
	_msg := _update.Message
	if _msg == nil {
		return false
	}
	if _msg.MessageID == wv.BaseIndex {
		return false
	}
	if _msg.From == nil {
		return false
	}
	if _msg.Chat == nil {
		return false
	}
	if _msg.Date == wv.BaseIndex {
		return false
	}
	if ws.IsEmpty(&_msg.Text) {
		return false
	}
	if _msg.Animation != nil {
		return false
	}
	if _msg.Document != nil {
		return false
	}
	if !ws.IsEmpty(&_msg.Caption) {
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
	if _msg.Audio != nil {
		return false
	}

	return true
}

func _checkGifMessage(_update *tg.Update) bool {

	return false
}

func _checkPhotoMessage(_update *tg.Update) bool {
	_msg := _update.Message
	if _msg == nil {
		return false
	}
	if _msg.MessageID == wv.BaseIndex {
		return false
	}
	//if _msg.From == nil {
	//	return false
	//}
	if _msg.Chat == nil {
		return false
	}
	if _msg.Date == wv.BaseIndex {
		return false
	}
	if ws.IsEmpty(&_msg.Text) {
		return false
	}
	if _msg.Animation != nil {
		return false
	}
	if _msg.Document != nil {
		return false
	}
	if !ws.IsEmpty(&_msg.Caption) {
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
	if _msg.Audio != nil {
		return false
	}

	return true
}
