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

var _alreadyRunning bool
var _settings *AppSettings

type AppSettings struct {
	botAPI     *tgbotapi.BotAPI
	router     *gin.Engine
	isGlobal   bool
	tObt       string
	port       string
	_apiObtain func(*AppSettings) bool
}

func GetSettings(_port string) *AppSettings {
	if _settings != nil {
		return GetExisting()
	}
	settings := AppSettings{}
	settings._setPort(&_port)
	settings._setRouter()
	_settings = &settings
	return _settings
}

func GetExisting() *AppSettings {
	return _settings
}

func (_s *AppSettings) InvalidateAPI() bool {
	if _s._apiObtain != nil {
		_s.isGlobal = !wotoSecurity.IsEmpty(&_s.tObt)
	} else {
		_s.isGlobal = false
	}
	return _s.isGlobal && _s._apiObtain(_s)
}

func (_s *AppSettings) SetAPI(_api *tgbotapi.BotAPI) {
	_s.botAPI = _api
}

func (_s *AppSettings) SetObt(_obt func(*AppSettings) bool) {
	_s._apiObtain = _obt
}

func (_s *AppSettings) SetTObt(_obt string) {
	_s.tObt = _obt
}

func (_s *AppSettings) _setPort(_port *string) {
	_s.port = *_port
}
func (_s *AppSettings) _setRouter() {
	_s.router = gin.New()
}

func (_s *AppSettings) GetAPI() *tgbotapi.BotAPI {
	return _s.botAPI
}

func (_s *AppSettings) IsGlobal() bool {
	return _s.isGlobal
}

func (_s *AppSettings) GetObt() string {
	return _s.tObt
}

func (_s *AppSettings) GetPort() string {
	return _s.port
}

func (_s *AppSettings) GetRouter() *gin.Engine {
	return _s.router
}
func App_enter() {
	_alreadyRunning = true
}

func App_exit() {
	_alreadyRunning = false
}

func IsRunning() bool {
	return _alreadyRunning
}
