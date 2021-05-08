// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoSudo

import (
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues/tgMessages/tgConst"
)

// the messages for wotoSudo plugin.
const (
	ADDED_TO_SUDO_LIST_M = "owo >>.<<\nUser " +
		tgConst.FIRST_NAME_MENTION +
		" has been added to my sudo list! -.-"
	ADDED_TO_SUDO_LIST_MPV = "owo damn you master >>.<<\nAnother trash " +
		tgConst.FIRST_NAME_MENTION +
		" has been added to my sudo list! -.-"

	ALREADY_SUDO_M = "But User " +
		tgConst.FIRST_NAME_MENTION +
		" is already in my sudo list! >.<"
	ALREADY_SUDO_MPV = "Idiot, User " +
		tgConst.FIRST_NAME_MENTION +
		" is already in my fucking sudo list! >.<"

	REMOVED_SUDO_M = "AwA ~.~\nUser " +
		tgConst.FIRST_NAME_MENTION +
		" is no longer in my sudo list! ^-^"
	REMOVED_SUDO_MPV = "AwA ~.~\n this trash " +
		tgConst.FIRST_NAME_MENTION +
		" is no longer in my bullshit sudo list! ^-^"

	NOT_IN_SUDO_LIST_M = "But User " +
		tgConst.FIRST_NAME_MENTION +
		" is not in my sodu list! *.*"
	NOT_IN_SUDO_LIST_MPV = "Open your goddamn eyes and see that " +
		tgConst.FIRST_NAME_MENTION +
		" is not in my sodu list! *.*"

	INVALID_ID_SUDO_M = "Id " + wotoValues.FORMAT_VALUE +
		" is not in correct format! ¯\\_(ツ)_/¯"
	INVALID_ID_SUDO_MPV = "You dumb! Id " + wotoValues.FORMAT_VALUE +
		" is not in correct format! (-.-)"
)

// the log messages for wotoSudo plugin.
const (
	ID_CANNOT_BE_ZERO = "In wotoSudo/handler.go: " +
		"ID cannot be zero!"
	ADD_SUDO_ERR = "Something happened in the " +
		"settings.AddSudo(), and result wasn't SUCCESS!"
	REM_SUDO_ERR = "Something happened in the " +
		"settings.RemSudo(), and result wasn't SUCCESS!"
)
