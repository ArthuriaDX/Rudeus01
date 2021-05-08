// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package tgConst

import (
	"strings"

	"github.com/ALiwoto/rudeus01/wotoPacks/wotoMD"
	ws "github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	USER_NAME          = "$WOTO_USER_NAME"
	BOT_USERNAME       = "$WOTO_BOT_USERNAME"
	FIRST_NAME         = "$WOTO_FIRST_NAME"
	USER_NAME_MENTION  = "$WOTO_USERNAME_MENTION"
	FIRST_NAME_MENTION = "$WOTO_FIRSTNAME_MENTION"
)

func ReplaceDyn(value *string, msg *tg.Message) {
	if ws.IsEmpty(value) || msg == nil {
		return
	}
	*value = wotoMD.GetNormal(*value).ToString()
	var tmp, md, userMd string

	md = wotoMD.GetNormal(USER_NAME).ToString()
	if strings.Contains(*value, md) {
		userMd = wotoMD.GetNormal(msg.From.UserName).ToString()
		tmp = strings.ReplaceAll(*value, md, userMd)
		*value = tmp
	}

	md = wotoMD.GetNormal(FIRST_NAME).ToString()
	if strings.Contains(*value, md) {
		userMd = wotoMD.GetNormal(msg.From.UserName).ToString()
		tmp = strings.ReplaceAll(*value, md, userMd)
		*value = tmp
	}

	md = wotoMD.GetNormal(USER_NAME_MENTION).ToString()
	if strings.Contains(*value, md) {
		// msg.From.ID doesn't need to be repaired at all,
		// and also msg.From.UserName will be repaired inside of the
		// function, so there is no need to repair for it.
		tmp = wotoMD.GetUserMention(msg.From.UserName, msg.From.ID).ToString()
		tmp = strings.ReplaceAll(*value, md, tmp)
		*value = tmp
	}

	md = wotoMD.GetNormal(FIRST_NAME_MENTION).ToString()
	if strings.Contains(*value, md) {
		// msg.From.ID doesn't need to be repaired at all,
		// and also msg.From.FirstName will be repaired inside of the
		// function, so there is no need to repair for it.
		tmp = wotoMD.GetUserMention(msg.From.FirstName, msg.From.ID).ToString()
		tmp = strings.ReplaceAll(*value, md, tmp)
		*value = tmp
	}
}

func ReplaceDynBot(value *string, api *tg.BotAPI) {
	var tmp string

	if strings.Contains(*value, BOT_USERNAME) {
		tmp = strings.ReplaceAll(*value, BOT_USERNAME, api.Self.UserName)
		*value = tmp
	}
}

func ReplaceFirstNameMention(v *string, firstName string, id int64) {
	*v = wotoMD.GetNormal(*v).ToString()
	var tmp, md string

	md = wotoMD.GetNormal(FIRST_NAME_MENTION).ToString()
	if strings.Contains(*v, md) {
		// id doesn't need to be repaired at all,
		// also firstName will be repaired inside of the
		// function, so there is no need to repair it.
		tmp = wotoMD.GetUserMention(firstName, id).ToString()
		tmp = strings.ReplaceAll(*v, md, tmp)
		*v = tmp
	}
}
