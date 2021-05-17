// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package appSettings

import (
	"log"
	"strconv"
	"strings"

	"github.com/ALiwoto/rudeus01/wotoPacks/interfaces"
	wa "github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/common"
	ws "github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (_s *AppSettings) SetAPI(_api *tg.BotAPI) {
	_s.botAPI = _api
}

func (_s *AppSettings) SetObt(_obt func(interfaces.WSettings) bool) {
	_s._apiObtain = _obt
}

func (_s *AppSettings) SetNextChild(_nextRunner func(string, interfaces.WSettings)) {
	_s._nextChild = _nextRunner
}

func (_s *AppSettings) SetURL(_url string) {
	if !ws.IsEmpty(&_url) {
		_s.url = _url
	}
}

func (_s *AppSettings) SetNextURL(_url string) {
	if !ws.IsEmpty(&_url) {
		_s.nextUrl = _url
	}
}

func (_s *AppSettings) SetTObt(_obt string) {
	_s.tObt = _obt
}

func (_s *AppSettings) SetIndex(_index int) {
	_s.index = _index
}

func (_s *AppSettings) SetSudoList(dbValue string) {
	strs := strings.Split(dbValue, wv.SEMICOLON)
	if strs == nil {
		log.Println(wv.SUDOLIST_SET_SETTINGS)
		return
	}

	_s.sudoList = make([]int64, len(strs))
	var err error
	for i, str := range strs {
		_s.sudoList[i], err = strconv.ParseInt(str, wv.BaseTen, wv.Base64Bit)
		if err != nil {
			log.Println(wv.SUDOLIST_SET_SETTINGS)
			continue
		}
	}
}

func (_s *AppSettings) SetMainSudo(id int64) {
	if _s.mainSudo != id {
		_s.mainSudo = id
	}
}

func (_s *AppSettings) AddSudo(id int64) wa.RESULT {
	if !_s.IsSudo(id) {
		re := _s.wClient.AddSudo(id)
		if re != wa.SUCCESS {
			log.Println(wv.SUDO_ADD_SETTINGS)
		}
		return re
	} else {
		return wa.CANCELED
	}
}

func (_s *AppSettings) RemSudo(id int64) wa.RESULT {
	if _s.IsSudo(id) {
		re := _s.wClient.RemSudo(id)
		if re != wa.SUCCESS {
			log.Println(wv.SUDO_ADD_SETTINGS)
		}
		return re
	} else {
		log.Println("This user is not a sudo!")
		return wa.CANCELED
	}
}

func (_s *AppSettings) SetWClient(_client interfaces.WClient) {
	_s.wClient = _client
}
