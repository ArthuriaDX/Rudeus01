// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package appSettings

import (
	"log"
	"net/http"

	"github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity"
)

func (_s *AppSettings) InvalidateAPI() bool {
	if _s._apiObtain != nil {
		_s.isGlobal = !wotoSecurity.IsEmpty(&_s.tObt)
		b := _s._apiObtain(_s)
		if !b {
			if _s._nextChild != nil {
				_s._nextChild(_s.url, _s)
			}
		}
		return b
	} else {
		_s.isGlobal = false
		return _s.isGlobal
	}
}

func App_enter() {
	_alreadyRunning = true
}

func App_exit() {
	_alreadyRunning = false
}

func (_s *AppSettings) RunNext() {
	if !wotoSecurity.IsEmpty(&_s.nextUrl) {
		go func() {
			_, _err := http.Get(_s.nextUrl)
			if _err != nil {
				log.Println(_err, _s.nextUrl)
			}
		}()
	}
}
