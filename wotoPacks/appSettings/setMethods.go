// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package appSettings

import (
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (_s *AppSettings) SetAPI(_api *tgbotapi.BotAPI) {
	_s.botAPI = _api
}

func (_s *AppSettings) SetObt(_obt func(*AppSettings) bool) {
	_s._apiObtain = _obt
}

func (_s *AppSettings) SetNextChild(_nextRunner func(string, *AppSettings)) {
	_s._nextChild = _nextRunner
}

func (_s *AppSettings) SetURL(_url string) {
	if !wotoSecurity.IsEmpty(&_url) {
		_s.url = _url
	}
}

func (_s *AppSettings) SetNextURL(_url string) {
	if !wotoSecurity.IsEmpty(&_url) {
		_s.nextUrl = _url
	}
}

func (_s *AppSettings) SetTObt(_obt string) {
	_s.tObt = _obt
}

func (_s *AppSettings) SetIndex(_index int) {
	_s.index = _index
}

func (_s *AppSettings) _setPort(_port *string) {
	_s.port = *_port
}
func (_s *AppSettings) _setRouter() {
	_s.router = gin.New()
}
