// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoValues

// the indexed constant values.
const (
	BaseIndex    = 0
	BaseOneIndex = 1
	BaseTimeOut  = 10
)

const (
	APP_PORT       = "PORT"
	GET_SLASH      = "/"
	HTTP_ADDRESS   = ":"
	TOKEN_KEY      = "TELEGRAM_TOKEN_BOT"
	RUDEUS_URL_KEY = "RUDEUS_URL"

	// INDEX_KEY is zero based, so it should
	// be started from 0.
	INDEX_KEY = "APP_INDEX"

	// TOTAL_INDEX_KEY is not zero based,
	// so it should NOT be start from zero.
	TOTAL_INDEX_KEY = "TOTAL_INDEX"
)

var DebugMode bool
