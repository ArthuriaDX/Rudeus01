// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoTranslate

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/wotoTranslate/wotoLang"
	ws "github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
)

// https://telegra.ph/Lang-Codes-03-19-3

// TranslateToMorse will translate the specified text value
// to the morse value
func Translate(text string) string {
	if ws.IsEmpty(&text) {
		return wv.EMPTY
	}
	lang := DetectLanguage(text)

	det := lang.Data.Detections
	if len(det) == wv.BaseIndex {
		return wv.EMPTY
	}

	fr := det[wv.BaseIndex].TheLang
	to := wotoLang.L_en

	return trTxt(fr, to, text)
}

func trTxt(fr, to, text string) string {
	urlG := fmt.Sprintf(gHostUrl, fr, to, text)
	resp, errH := http.Get(urlG)
	if errH != nil {
		log.Println(errH)
		return wv.EMPTY
	}

	defer resp.Body.Close()

	b, errB := ioutil.ReadAll(resp.Body)
	if errB != nil {
		log.Println(errB)
		return wv.EMPTY
	}

	return string(b)
}
