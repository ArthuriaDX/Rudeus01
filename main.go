// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/wotoChilds"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv(wv.APP_PORT)
	if wotoSecurity.IsEmpty(&port) {
		log.Fatal(wv.PORT_ERROR)
	}
	settings := appSettings.GetSettings(port)
	runApp(settings)
}

func runApp(_settings *appSettings.AppSettings) {
	router := _settings.GetRouter()
	router.Use(gin.Logger())
	router.GET(wv.GET_SLASH, answerClient)
	_ = router.Run(wv.HTTP_ADDRESS + _settings.GetPort())
}

func answerClient(c *gin.Context) {
	if !appSettings.IsRunning() {
		appSettings.App_enter()
	} else {
		_s := appSettings.GetExisting()
		_s.RunNext()
		c.String(http.StatusAccepted, wv.FORMAT_VALUE, wv.ALREADY_RUNNING)
		return
	}
	wv.DebugMode = true
	_settings := appSettings.GetExisting()
	_token := os.Getenv(wv.TOKEN_KEY)
	if wotoSecurity.IsEmpty(&_token) {
		log.Fatal(wv.INVALID_API)
	}
	_settings.SetTObt(_token)
	_settings.SetObt(wotoChilds.WObtainTC)
	_settings.SetNextChild(wotoChilds.RunNextChild)
	if !_settings.InvalidateAPI() {
		c.String(http.StatusForbidden, wv.FORMAT_VALUE, wv.INVALID_ENGINE)
		log.Println(wv.INVALID_ENGINE)
		os.Exit(wv.BaseIndex)
	}
	result, client := wotoActions.GenerateClient()
	if result != wotoActions.SUCCESS {
		log.Fatal(wv.WOTO_CLIENT_ERROR)
	}
	wotoActions.RunBot(_settings, client)
}
