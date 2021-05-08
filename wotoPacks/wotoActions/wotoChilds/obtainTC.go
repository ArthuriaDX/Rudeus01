// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoChilds

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/ALiwoto/rudeus01/wotoPacks/interfaces"
	ws "github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
)

const (
	_WORKER           = "https://frosty-surf-5435.saogame.workers.dev/"
	_REQUEST_VALUE    = "request_value"
	_TOKEN_VALUE      = "<get_token>"
	_APPLICATION_FORM = "application/json"
)

func WObtainTC(_settings interfaces.WSettings) bool {

	values := map[string]string{_REQUEST_VALUE: _TOKEN_VALUE}
	jdata, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}
	indexStr := os.Getenv(wotoValues.INDEX_KEY)
	if ws.IsEmpty(&indexStr) {
		_settings.SetIndex(wotoValues.BaseIndex)
		return true
	}
	url := os.Getenv(wotoValues.RUDEUS_URL_KEY)
	if !ws.IsEmpty(&url) {
		_settings.SetURL(url)
	} else {
		log.Println()
	}

	index, _err := strconv.Atoi(indexStr)

	if _err != nil {
		// do NOT use fatal here, it will cuase trouble for bot
		log.Println(_err)
	}

	_settings.SetIndex(index)
	resp, err := http.Post(_WORKER, _APPLICATION_FORM, bytes.NewBuffer(jdata))

	if err != nil {
		// do NOT use fatal here, it will cuase trouble for bot
		log.Println(err)
		return false
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// do NOT use fatal here, it will cuase trouble for bot
		log.Println(err)
		return false
	}
	str := string(body)
	myInt, _err := strconv.Atoi(str)
	if _err != nil {
		return false
	}
	return myInt == index
}
