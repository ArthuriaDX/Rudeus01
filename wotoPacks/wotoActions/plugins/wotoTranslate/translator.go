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
func Translate(fr, to, text string) []string {
	if ws.IsEmpty(&text) {
		return nil
	}

	text = trGoogle(fr, to, text)
	s := parseGData(text)

	return s
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

func parseGData(text string) []string {
	test := ws.Split(text, "[", "]")
	original := make([]string, wv.BaseIndex)
	accepted := func(v string) bool {
		if v == "\\n" {
			return false
		}
		if v == "\\n," {
			return false
		}
		if strings.Contains(v, "af.httprm") {
			return false
		}
		if strings.Contains(v, "\"e\",4,") {
			return false
		}
		if strings.HasPrefix(v, "null,") {
			tmpN := strings.ReplaceAll(v, "null,", wv.EMPTY)
			_, errN := strconv.Atoi(tmpN)
			if errN == nil {
				return false
			}
		}
		if strings.HasPrefix(v, "\"di\"") {
			return false
		}
		tmp := strings.ReplaceAll(v, " ", wv.EMPTY)
		tmp = strings.ReplaceAll(tmp, ",", wv.EMPTY)
		if len(tmp) == wv.BaseIndex {
			return false
		}
		if tmp == ")" || tmp == "}'" {
			return false
		}
		if !strings.Contains(tmp, "null") && !strings.Contains(tmp, "\"") {
			return false
		}
		if strings.Contains(v, "\"wrb.fr\",\"MkEWBc\"") {
			return false
		}
		tmp = strings.ReplaceAll(tmp, "\n", wv.EMPTY)
		_, errI := strconv.Atoi(tmp)
		return errI != nil
	}

	for _, s := range test {
		if accepted(s) {
			original = append(original, s)
		}
	}

	return original
}

func GetOpening(text string) string {
	finalStr := wv.EMPTY
	// ,[[\"
	// \",[\"
	c1 := false
	c2 := false
	c3 := false
	c4 := false
	after := false

	cA1 := false
	cA2 := false
	cA3 := false
	cA4 := false
	cA5 := false
	ended := false
	invalidAll := func() {
		if c1 {
			c1 = false
		}
		if c2 {
			c2 = false
		}
		if c3 {
			c3 = false
		}
		if c4 {
			c4 = false
		}
	}
	invalidAfter := func() {
		if cA1 {
			cA1 = false
		}
		if cA2 {
			cA2 = false
		}
		if cA3 {
			cA3 = false
		}
		if cA4 {
			cA4 = false
		}
		if cA5 {
			cA5 = false
		}
	}

	for _, ch := range text {
		if after {
			if ended {
				finalStr += string(ch)
				continue
			}
			if cA5 {
				if ch == '"' {
					ended = true
				} else {
					invalidAfter()
				}
				continue
			}
			if cA4 {
				if ch == '\\' {
					cA5 = true
				} else {
					invalidAfter()
				}
				continue
			}
			if cA3 {
				if ch == '[' {
					cA4 = true
				} else {
					invalidAfter()
				}
				continue
			}
			if cA2 {
				if ch == ',' {
					cA3 = true
				} else {
					invalidAfter()
				}
				continue
			}
			if cA1 {
				if ch == '"' {
					cA2 = true
				} else {
					invalidAfter()
				}
				continue
			}

			if ch == '\\' {
				cA1 = true
			}

			continue
		}

		if c4 {
			if ch == '"' {
				after = true
			} else {
				invalidAll()
			}
			continue
		}

		if c3 {
			if ch == '\\' {
				c4 = true
			} else {
				invalidAll()
			}
			continue
		}

		if c2 {
			if ch == '[' {
				c3 = true
			} else {
				invalidAll()
			}
			continue
		}

		if c1 {
			if ch == '[' {
				c2 = true
			} else {
				invalidAll()
			}
			continue
		}

		if ch == ',' {
			c1 = true
			continue
		}
	}
	return finalStr
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
	if strings.Contains(text, "[") {
		text = strings.ReplaceAll(text, "[", "((((")
	}
	if strings.Contains(text, "]") {
		text = strings.ReplaceAll(text, "]", "))))")
	}
	if strings.Contains(text, "*") {
		text = strings.ReplaceAll(text, "*", wv.EMPTY)
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
	return fReqGValue1 + url.PathEscape(text) +
		fReqGValue2 + fr +
		fReqGValue3 + to + fReqGValue4
}
