// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package appSettings

import (
	"github.com/ALiwoto/rudeus01/wotoPacks/interfaces"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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

// GetIndex will give you the index of the current child.
// please notice that this value is zero-based.
func (_s *AppSettings) GetIndex() int {
	return _s.index
}

// GetTotalIndex will give you the total index of childs.
// please notice that this value is NOT zero-based.
func (_s *AppSettings) GetTotalIndex() int {
	return _s.totalIndex
}

func (_s *AppSettings) GetWClient() interfaces.WClient {
	return _s.wClient
}

func (_s *AppSettings) GetSudoList() []int64 {
	return _s.sudoList
}

func (_s *AppSettings) IsSudo(id int64) bool {
	if _s.sudoList == nil {
		return false
	}

	for _, current := range _s.sudoList {
		if id == current {
			return true
		}
	}

	return false
}

func (_s *AppSettings) IsMainSudo(id int64) bool {
	if _s.mainSudo == wotoValues.BaseIndex {
		return false
	}
	return _s.mainSudo == id
}
