// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoPat

import (
	"time"

	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
)

func getRandomPat() (string, patType) {
	if patMap == nil || len(*patMap) == wv.BaseIndex {
		return wv.EMPTY, NONE
	}

	tmp := time.Now().Nanosecond() % len(*patMap)
	id := (*patMap)[tmp][wv.PAT_ID_KEY].(string)
	t := (*patMap)[tmp][wv.PAT_TYPE_KEY].(int32)
	return id, patType(t)
}

func getRandomHPat() (string, patType) {
	if patHMap == nil || len(*patHMap) == wv.BaseIndex {
		return wv.EMPTY, NONE
	}

	tmp := time.Now().Nanosecond() % len(*patHMap)
	id := (*patHMap)[tmp][wv.PAT_ID_KEY].(string)
	t := (*patHMap)[tmp][wv.PAT_TYPE_KEY].(int32)
	return id, patType(t)
}
