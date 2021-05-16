// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoTranslate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	ws "github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
)

// https://telegra.ph/Lang-Codes-03-19-3

// Translate will translate the specified text value
// tp english.
func Translate(fr, to, text string) *WotoTr {
	if ws.IsEmpty(&text) {
		return nil
	}

	text = trGoogle(fr, to, text)
	w := WotoTr{
		OriginalText: text,
		From:         fr,
		To:           to,
	}

	return parseGData(&w)
}

func TrGnuTxt(fr, to, text string) string {
	urlG := fmt.Sprintf(gnuHostUrl, fr, to, text)
	resp, errH := http.Get(urlG)
	if errH != nil {
		log.Println(errH)
		return errTrGnuText
	}

	defer resp.Body.Close()

	b, errB := ioutil.ReadAll(resp.Body)
	if errB != nil {
		log.Println(errB)
		return errTrGnuText
	}

	var g gnuTranslate
	errJ := json.Unmarshal(b, &g)
	if errJ != nil {
		log.Println(errJ)
		return errTrGnuText
	}

	if ws.IsEmpty(&g.Err) {
		log.Println(g.Err)
		return errTrGnuText
	}

	return g.Result
}

func parseGData(wTr *WotoTr) *WotoTr {
	text := wTr.OriginalText
	test := ws.Split(text, wv.BracketOpen, wv.Bracketclose)
	original := make([]string, wv.BaseIndex)
	accepted := func(v string) bool {
		if v == NonEscapeN {
			return false
		}
		if v == NonEscapeNV {
			return false
		}
		if strings.Contains(v, HttpRm) {
			return false
		}
		if strings.Contains(v, E4Value) {
			return false
		}
		if strings.HasPrefix(v, NullCValue) {
			tmpN := strings.ReplaceAll(v, NullCValue, wv.EMPTY)
			_, errN := strconv.Atoi(tmpN)
			if errN == nil {
				return false
			}
		}
		if strings.HasPrefix(v, DiValue) {
			return false
		}
		tmp := strings.ReplaceAll(v, wv.SPACE_VALUE, wv.EMPTY)
		tmp = strings.ReplaceAll(tmp, wv.CAMA, wv.EMPTY)
		if len(tmp) == wv.BaseIndex {
			return false
		}
		if tmp == wv.ParaClose || tmp == AkCloseQ {
			return false
		}
		if !strings.Contains(tmp, wv.NullStr) &&
			!strings.Contains(tmp, wv.DoubleQ) {
			return false
		}
		if strings.Contains(v, WrbFr) {
			return false
		}
		tmp = strings.ReplaceAll(tmp, wv.N_ESCAPE, wv.EMPTY)
		_, errI := strconv.Atoi(tmp)
		return errI != nil
	}

	for _, s := range test {
		if accepted(s) {
			original = append(original, s)
		}
	}

	strs := parseGparams(original, wTr)

	return arrangeParams(strs, wTr)
}

func parseGparams(value []string, wTr *WotoTr) []string {
	//null,
	//null,
	// \"ja\"
	// \n,null,
	// null,\"Konnichiwa. Ohayou Minna\",null,null,null,
	// \"konnichiwa。\",
	// \"konnichiwa。\",\"こんにちは。\"
	// \"Ohayou Minna\",
	// \"Ohayou Minna\",\"みんなおはよう\"
	// \n,\"ja\",1,\"en\",
	// \"konnichiwa. ohayou minna \",\"en\",\"ja\",true
	// \n",null,null,null,"generic"
	if wTr.Road == nil {
		wTr.Road = make(map[int]bool)
	}

	index := wv.BaseIndex

	for _, c := range wTr.OriginalText {
		if string(c) == wv.N_ESCAPE {
			wTr.Road[index] = false
		}
		if string(c) == wv.Point {
			wTr.Road[index] = true
		}
		index++
	}
	tmp := strings.Join(value, wv.DY_WOTO_TEXT)
	tmp = strings.ReplaceAll(tmp, NullN, wv.EMPTY)
	tmp = strings.ReplaceAll(tmp, NullCValueR, wv.EMPTY)
	tmp = strings.ReplaceAll(tmp, GenericStr, wv.EMPTY)
	tmp = strings.ReplaceAll(tmp, NullCValue, wv.EMPTY)
	tmp = strings.ReplaceAll(tmp, NeQ, wv.EMPTY)
	strs := strings.Split(tmp, wv.DY_WOTO_TEXT)
	final := make([]string, wv.BaseIndex)
	strMap := make(map[string]bool)
	lastStr := wv.EMPTY
	for _, current := range strs {
		tmp = current
		if current == lastStr {
			continue
		}
		if strings.HasPrefix(current, DoubleQS) {
			current = strings.TrimPrefix(current, DoubleQS)
		} else {
			lastStr = wv.EMPTY
			continue
		}

		if strings.Contains(current, MiddleWave) {
			current = strings.Split(current, MiddleWave)[wv.BaseOneIndex]
		}

		if strings.HasSuffix(current, DoubleQSP) {
			// optional
			current = strings.TrimSuffix(current, DoubleQSP)

			if strMap[current] {
				continue
			} else {
				strMap[current] = true
			}

			final = append(final, current)
			lastStr = tmp
			continue
		}

		if strings.HasSuffix(current, DoubleQS) {
			current = strings.TrimSuffix(current, DoubleQS)
		} else {
			lastStr = wv.EMPTY
			continue
		}

		if strMap[current] {
			continue
		} else {
			strMap[current] = true
		}

		lastStr = tmp
		log.Println(current)
		log.Println(strMap)
		final = append(final, current)
	}

	return final
}

func arrangeParams(values []string, wTr *WotoTr) *WotoTr {
	index := wv.BaseIndex
	for _, current := range values {
		if strings.Contains(current, WrongNessOpen) {
			wTr.HasWrongNess = true
			current = strings.ReplaceAll(current, WrongNessOpen, wv.EMPTY)
			current = strings.ReplaceAll(current, WrongNessClose, wv.EMPTY)
			current = strings.TrimPrefix(current, wv.BackSlash)
			wTr.CorrectedValue = current
		} else {
			if wTr.Road != nil {
				if !wTr.Road[index] {
					current += wv.N_ESCAPE
				} else {
					current = strings.TrimPrefix(current, wv.N_ESCAPE)
					current = strings.TrimSuffix(current, wv.N_ESCAPE)
				}
			}
			wTr.TranslatedText += current
		}
	}
	// TODO!
	//wTr.TranslatedText = strings.Join(values, wv.N_ESCAPE)
	return wTr
}

// trGAuto will translate the `text`
// from "auto" to `to` string and it will return the
// velue.
func AtrGAuto(to, text string) []string {
	return nil
}

func trGoogle(fr, to, text string) string {
	body := strings.NewReader(googleFQ(fr, to, purify(text)))
	req, err := http.NewRequest(requestType, gHostUrl, body)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set(userAgentGKey, userAgentGValue)
	req.Header.Set(acceptGKey, acceptGValue)
	req.Header.Set(acceptLanguageGKey, acceptLanguageGValue)
	req.Header.Set(refererGKey, refererGValue)
	req.Header.Set(xSameDomainGKey, xSameDomainGValue)
	req.Header.Set(xGoogBatchExecuteBgrGKey, xGoogBatchExecuteBgrGValue)
	req.Header.Set(contentTypeGKey, contentTypeGValue)
	req.Header.Set(originGKey, originGValue)
	req.Header.Set(gDNTGKey, gDNTGValue)
	req.Header.Set(connectionGKey, connectionGValue)
	req.Header.Set(cookieGKey, cookieGValue)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	b, errB := ioutil.ReadAll(resp.Body)
	if errB != nil {
		log.Fatal(errB)
	}

	return string(b)
}

func purify(text string) string {
	if strings.Contains(text, wv.BracketOpen) {
		text = strings.ReplaceAll(text, wv.BracketOpen, wv.ParaOpen)
	}
	if strings.Contains(text, wv.Bracketclose) {
		text = strings.ReplaceAll(text, wv.Bracketclose, wv.ParaClose)
	}
	if strings.Contains(text, wv.Star) {
		text = strings.ReplaceAll(text, wv.Star, wv.EMPTY)
	}
	if strings.Contains(text, wv.N_ESCAPE) {
		text = strings.ReplaceAll(text, wv.N_ESCAPE, wv.SPACE_VALUE)
	}
	if strings.Contains(text, wv.R_ESCAPE) {
		text = strings.ReplaceAll(text, wv.R_ESCAPE, wv.SPACE_VALUE)
	}
	if strings.Contains(text, wv.DoubleQ) {
		text = strings.ReplaceAll(text, wv.DoubleQ, wv.DoubleQJ)
	}
	return text
}

func googleFQ(fr, to, text string) string {
	//sUrl := url.PathEscape(text)
	//tUrl, _ := url.Parse(text)
	//sUrl := tUrl.String()
	//return "f.req=%5B%5B%5B%22MkEWBc%22%2C%22%5B%5B%5C%22How%20are%20you%5C%22%2C%5C%22auto%5C%22%2C%5C%22fa%5C%22%2Ctrue%5D%2C%5Bnull%5D%5D%22%2Cnull%2C%22generic%22%5D%5D%5D&"
	// [[["MkEWBc","[[\"Hello\",\"auto\",\"fa\",true],[null]]",null,"generic"]]]&
	//return "[[[\"MkEWBc\", \"[[\"Hello\",\"auto\",\"fa\",true],[null]]\",null,\"generic\"]]]&\""
	return fReqGValue1 + url.QueryEscape(text) +
		fReqGValue2 + fr +
		fReqGValue3 + to + fReqGValue4
}
