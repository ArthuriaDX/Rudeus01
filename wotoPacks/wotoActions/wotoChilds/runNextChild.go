// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoChilds

import (
	"strconv"
	"strings"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
)

const (
	_ZERO_PREFEX = "0"
	_LENGTH      = 2
)

// RunNextChild will run the next child. please consider that the url
// is the current child url, not the next one.
func RunNextChild(url string, settings *appSettings.AppSettings) {
	if wotoSecurity.IsEmpty(&url) || settings == nil {
		return
	}
	var next int
	index := settings.GetIndex()
	total := settings.GetTotalIndex()
	differ := total - wotoValues.BaseOneIndex
	if index == differ {
		next = wotoValues.BaseIndex
	} else {
		next = index + wotoValues.BaseOneIndex
	}
	next++
	indexStr := strconv.Itoa(index)
	nextStr := strconv.Itoa(next)
	if len(nextStr) != _LENGTH {
		nextStr = _ZERO_PREFEX + nextStr
	}

	finalURL := strings.ReplaceAll(url, indexStr, nextStr)
	if !wotoSecurity.IsEmpty(&finalURL) {
		settings.SetNextURL(finalURL)
	}
	settings.RunNext()
}
