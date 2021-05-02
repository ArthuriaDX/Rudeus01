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
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv(wotoValues.APP_PORT)
	if wotoSecurity.IsEmpty(&port) {
		log.Fatal(wotoValues.PORT_ERROR)
	}
	settings := appSettings.GetSettings(port)
	runApp(settings)
}

func runApp(_settings *appSettings.AppSettings) {
	router := _settings.GetRouter()
	router.Use(gin.Logger())
	router.GET(wotoValues.GET_SLASH, answerClient)
	_ = router.Run(wotoValues.HTTP_ADDRESS + _settings.GetPort())
}

func answerClient(c *gin.Context) {
	if !appSettings.IsRunning() {
		appSettings.App_enter()
	} else {
		c.String(http.StatusAccepted, "%v", wotoValues.ALREADY_RUNNING)
		return
	}
	wotoValues.DebugMode = true
	_settings := appSettings.GetExisting()
	_token := os.Getenv(wotoValues.TOKEN_KEY)
	if wotoSecurity.IsEmpty(&_token) {
		log.Fatal(wotoValues.INVALID_API)
	}
	_settings.SetTObt(_token)
	_settings.SetObt(wotoActions.WObtainTC)
	if !_settings.InvalidateAPI() {
		c.String(http.StatusForbidden, "%v", wotoValues.INVALID_ENGINE)
		log.Println(wotoValues.INVALID_ENGINE)
		os.Exit(wotoValues.BaseIndex)
	}
	result, client := wotoActions.GenerateClient()
	if result != wotoActions.SUCCESS {
		log.Fatal(wotoValues.WOTO_CLIENT_ERROR)
	}
	wotoActions.RunBot(_settings, client)
}
