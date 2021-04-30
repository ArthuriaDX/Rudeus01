package wotoActions

import (
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

	return false
}

func _checkGifMessage(_update *tgbotapi.Update) bool {

	return false
}
