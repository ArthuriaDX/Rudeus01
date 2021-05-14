// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoUD

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	ws "github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
)

// the main structure for interacting with urban dictionary
// public API.
type UrbanDictionary struct {
	Def        string   `json:"definition"`
	PermLink   string   `json:"permalink"`
	ThumbsUp   int      `json:"thumbs_up"`
	Voices     []string `json:"sound_urls"`
	Author     string   `json:"author"`
	Word       string   `json:"word"`
	DefID      int      `json:"defid"`
	Vote       string   `json:"current_vote"`
	Date       string   `json:"written_on"`
	Example    string   `json:"example"`
	ThumbsDown int      `json:"thumbs_down"`
}

// UrbanCollection is a collection (array) of Urban dictionary
// values.
type UrbanCollection struct {
	List []UrbanDictionary `json:"list"`
}

//type similarity struct {
//	array []simArray
//}

//type simArray struct {
//	value string
//}

// Define, will define the word, using
// urban dictionary public API, version 0? :/
// It will return a pointer to an urban collection.
func Define(word string) (ud *UrbanCollection, err error) {
	if ws.IsEmpty(&word) {
		return nil, errors.New(emptyWordErr)
	}

	word = strings.ToLower(word)
	word = strings.TrimSpace(word)

	resp, httpErr := http.Get(hostUrl + url.PathEscape(word))
	if httpErr != nil {
		return nil, httpErr
	}
	if resp == nil || resp.Body == nil {
		return nil, errors.New(nilRespErr)
	}

	defer resp.Body.Close()

	jData, dError := ioutil.ReadAll(resp.Body)
	if dError != nil {
		return nil, dError
	}
	if jData == nil || len(jData) == wv.BaseIndex {
		return nil, errors.New(jDataErr)
	}

	udValue := UrbanCollection{}

	//log.Println(string(jData))
	jErr := json.Unmarshal(jData, &udValue)

	if jErr != nil {
		return nil, jErr
	}

	return &udValue, nil
}
