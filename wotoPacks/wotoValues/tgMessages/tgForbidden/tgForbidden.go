// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package tgForbidden

import "strings"

// these values are values which will be received from telegram as
// `Forbidden` errors.
const (
	FORBIDDEN  = "Forbidden"
	START_CHAT = "Forbidden: bot can't initiate conversation with a user"
)

type ForbiddenError struct {
	_value string
}

func GetForbidden(err error) (*ForbiddenError, bool) {
	str := err.Error()
	if strings.Contains(str, FORBIDDEN) {
		fr := ForbiddenError{
			_value: str,
		}
		return &fr, true
	}
	return nil, false
}

func (_f *ForbiddenError) IsPrivateMessageError() bool {
	return strings.Contains(_f._value, START_CHAT)
}
