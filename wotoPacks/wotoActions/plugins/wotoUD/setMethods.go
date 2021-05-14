// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoUD

import tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"

//---------------------------------------------------------

//---------------------------------------------------------
func (o *udOrigin) setCollection(c *UrbanCollection) {
	if o.collection != c {
		o.collection = c
	}
}

func (o *udOrigin) setButtons(b *tg.InlineKeyboardMarkup) {
	if o.keyboard != b {
		o.keyboard = b
	}
}

//---------------------------------------------------------
