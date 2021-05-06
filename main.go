// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/interfaces"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/common"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/wotoChilds"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoDB"
	ws "github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	"github.com/gin-gonic/gin"
)

func main() {
	resp, err1 := http.Get("https://api.urbandictionary.com/v0/define?term=mamad")
	if err1 != nil {
		log.Fatal(err1)
	}

	str, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Fatal(err2)
	}
	log.Println(string(str))

	port := os.Getenv(wv.APP_PORT)
	if ws.IsEmpty(&port) {
		log.Fatal(wv.PORT_ERROR)
	}
	settings := appSettings.GetSettings(port)
	runApp(settings)
}

func runApp(_settings interfaces.WSettings) {
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
		_s.SetObt(wotoChilds.WObtainTC)
		if !_s.InvalidateAPI() {
			c.String(http.StatusForbidden, wv.FORMAT_VALUE, wv.INVALID_ENGINE)
			log.Fatal(wv.INVALID_ENGINE)
		}
		c.String(http.StatusAccepted, wv.FORMAT_VALUE, wv.ALREADY_RUNNING)
		return
	}
	wv.DebugMode = true
	_settings := appSettings.GetExisting()
	_token := os.Getenv(wv.TOKEN_KEY)
	if ws.IsEmpty(&_token) {
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
	result, client := wotoDB.GenerateClient()
	if result != common.SUCCESS {
		log.Fatal(wv.WOTO_CLIENT_ERROR)
	}

	_settings.SetWClient(client)
	go tickChecker()
	wotoActions.RunBot(_settings, client)
}

// the tick checker for checking the authority of current engine.
func tickChecker() {
	_s := appSettings.GetExisting()
	for {
		time.Sleep(wv.TICK_INTERVAL)
		if !wotoChilds.WObtainTC(_s) {
			log.Fatal(wv.INVALID_ENGINE)
		}
	}
}
