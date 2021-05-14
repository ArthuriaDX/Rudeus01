// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoUD

import (
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type udQuery struct {
	next     bool // true if it's the next button
	previous bool // true if it's the previous button
	voice    bool // true if it's the voice button
	// sound    bool      // true if it's the voice button (TODO)
	origin *udOrigin // the origin of the ud API.

}

type udOrigin struct {
	currentPage uint8                    // must be between 1 and 10.
	word        string                   // the word used in urban dictionary
	id          int                      // the message id
	uId         int64                    // the user id (owner of the ud)
	collection  *UrbanCollection         // the collection
	keyboard    *tg.InlineKeyboardMarkup //the buttons
	uniqueId    string                   // look, it should be the date
}

// getNewOrigin will give you a new origin.
func getNewOrigin(w string, mId int, userId int64, unique string) (origin *udOrigin) {
	return &udOrigin{
		currentPage: wv.BaseOneIndex,
		word:        w,
		id:          mId,
		uId:         userId,
		uniqueId:    unique,
	}
}

func getNextUdQuery(origin *udOrigin) *udQuery {
	return getUdQuery(origin, true)
}

func getPreviousUdQuery(origin *udOrigin) *udQuery {
	return getUdQuery(origin, false)
}

func getVoiceUdQuery(origin *udOrigin) *udQuery {
	return &udQuery{
		next:     false,
		previous: false,
		voice:    true,
		origin:   origin,
	}
}

func getUdQuery(orig *udOrigin, next bool) *udQuery {
	return &udQuery{
		next:     next,
		previous: !next,
		origin:   orig,
	}
}
