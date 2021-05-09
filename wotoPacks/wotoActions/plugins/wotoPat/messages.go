// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoPat

import (
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues/tgMessages/tgConst"
)

// messages used in wotoPat plugin.
const (
	SEND_HENTAI_SUDO_ONLY = "Owo " +
		tgConst.FIRST_NAME_MENTION +
		", ~-~ Only sudo people can use hentai flag!"
	SEND_HENTAI_PV_ONLY = "#-# " +
		tgConst.FIRST_NAME_MENTION +
		", Sending lewd things in" +
		" groups is prohibited!\nConsider adding --pv flag :D"
	NO_HPAT = "UmU! " +
		tgConst.FIRST_NAME_MENTION +
		", it seems there is no hentai pat in the database! >-<"
	NO_PAT = "UmU! " +
		tgConst.FIRST_NAME_MENTION +
		", it seems there is no pat in the database! >.<"
	PAT_EXISTS = "Nyaaaa! pat =>` " +
		wotoValues.FORMAT_VALUE +
		"` , already exists in the database! -.-"
	PAT_DOESNT_EXIST = "Nyaaaa! pat =>` " +
		wotoValues.FORMAT_VALUE +
		"` , doesn't exist in the database! *-*"
	HENTAI_PAT_ADDED   = "Hentai pat has been added to database-nano~"
	PAT_ADDED          = "Pat has been added to database meow~☆"
	HENTAI_PAT_REMOVED = "Hentai pat has been removed from database-nano~"
	PAT_REMOVED        = "Pat has been removed from database meow~☆"
)
