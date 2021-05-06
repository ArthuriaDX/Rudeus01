// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoChilds

import (
	"strconv"
	"strings"

	"github.com/ALiwoto/rudeus01/wotoPacks/interfaces"
	ws "github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
)

const (
	_LENGTH = 2
)

// RunNextChild will run the next child. please consider that the url
// is the current child url, not the next one.
func RunNextChild(url string, settings interfaces.WSettings) {
	if ws.IsEmpty(&url) || settings == nil {
		return
	}
	var next int
	index := settings.GetIndex()
	total := settings.GetTotalIndex()
	differ := total - wv.BaseOneIndex
	if index == differ {
		next = wv.BaseIndex
	} else {
		next = index + wv.BaseOneIndex
	}
	next++
	indexStr := strconv.Itoa(index)
	nextStr := strconv.Itoa(next)
	if len(nextStr) != _LENGTH {
		nextStr = wv.BaseIndexStr + nextStr
	}

	finalURL := strings.ReplaceAll(url, indexStr, nextStr)
	if !ws.IsEmpty(&finalURL) {
		settings.SetNextURL(finalURL)
	}
	settings.RunNext()
}
