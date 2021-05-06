// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package interfaces

import (
	"github.com/gin-gonic/gin"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Go encourages composition over inheritance, using simple, often one-method interfaces ...
// that serve as clean, comprehensible boundaries between components.
// 	—Rob Pike,
//		“Go at Google: Language Design in the
//		Service of Software Engineering”
//		(see talks.golang.org/ 2012/splash.article)

type WSettings interface {
	RunNext()
	GetAPI() *tg.BotAPI
	IsGlobal() bool
	GetObt() string
	GetPort() string
	GetRouter() *gin.Engine
	GetIndex() int
	GetTotalIndex() int
	GetWClient() WClient
	SetURL(_url string)
	SetAPI(_api *tg.BotAPI)
	SetObt(_obt func(WSettings) bool)
	SetTObt(_obt string)
	SetNextURL(_url string)
	SetIndex(_index int)
	SendSudo(str string)
	InvalidateAPI() bool
	SetNextChild(_nextRunner func(string, WSettings))
	SetWClient(_client WClient)
}
