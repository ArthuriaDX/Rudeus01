// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package tgMessages

import "strings"

// these values are values which will be received from telegram as
// `Forbidden` errors.
const (
	FORBIDDEN  = "Forbidden"
	START_CHAT = "Forbidden: bot can't initiate conversation with a user"
)

func IsForbidden(err error) bool {
	str := err.Error()
	if strings.Contains(str, START_CHAT) {
		return true
	}
	if strings.Contains(str, FORBIDDEN) {
		return true
	}
	return false
}
