// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package main

import (
	"log"
	"os"
	"time"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/interfaces"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/wotoChilds"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoDB"
	ws "github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
)

func main() {
	settings := appSettings.GetSettings()
	runApp(settings)
}

func runApp(_settings interfaces.WSettings) {
	answerClient()
}

func answerClient() {
	if !appSettings.IsRunning() {
		//go c.String(http.StatusAccepted, wv.FORMAT_VALUE, wv.BaseTen)
		appSettings.App_enter()
	} else {
		_s := appSettings.GetExisting()
		_s.RunNext()
		_s.SetObt(wotoChilds.WObtainTC)
		if !_s.InvalidateAPI() {
			//c.String(http.StatusForbidden, wv.FORMAT_VALUE, wv.INVALID_ENGINE)
			log.Fatal(wv.INVALID_ENGINE)
		}
		//c.String(http.StatusAccepted, wv.FORMAT_VALUE, wv.ALREADY_RUNNING)
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
		//c.String(http.StatusForbidden, wv.FORMAT_VALUE, wv.INVALID_ENGINE)
		log.Println(wv.INVALID_ENGINE)
		os.Exit(wv.BaseIndex)
	}

	wotoDB.GenerateClients(_settings)
	//result, client := wotoDB.GenerateClient(1)
	//if result != common.SUCCESS {
	//	log.Fatal(wv.WOTO_CLIENT_ERROR)
	//}

	//_settings.SetWClient(client)
	go tickChecker()
	wotoActions.RunBot(_settings)
}

// the tick checker for checking the authority of current engine.
func tickChecker() {
	_s := appSettings.GetExisting()
	for {
		time.Sleep(wv.TICK_INTERVAL)
		if !wotoChilds.WObtainTC(_s) {
			_s.RunNext()
			log.Fatal(wv.INVALID_ENGINE)
		}
	}
}
