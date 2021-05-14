// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoUD

import (
	"fmt"

	"github.com/ALiwoto/rudeus01/wotoPacks/wotoMD"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
)

//---------------------------------------------------------

func (ud *udQuery) GetId() int {
	return ud.origin.id
}

func (ud *udQuery) GetUniqueId() string {
	return ud.origin.uniqueId
}

//---------------------------------------------------------

// GetDef will give you the definition in the specified page,
// using wotoMD.
// WARNING: page arg is zero indexed.
func (c *UrbanCollection) GetDef(page uint8) string {
	if len(c.List) == wv.BaseIndex {
		// TODO: return not found or something.
		return listNAvailble
	}

	//text = wotoMD.GetNormal(ud.List[0].Def).ToString()
	normal := wotoMD.GetNormal(c.List[page].Def + wv.N_ESCAPE + wv.N_ESCAPE)

	ex := wotoMD.GetItalic(wv.N_ESCAPE + wv.N_ESCAPE + c.List[page].Example)
	actual := page + wv.BaseOneIndex
	end := fmt.Sprintf(footnote, actual, len(c.List))
	final := normal.Append(ex).Append(wotoMD.GetNormal(end))
	return final.ToString()
}

// Found will give you a boolean indicating whether the word
// in the specified page has at least one correct voice url or not.
func (c *UrbanCollection) HasVoice(page uint8) bool {
	if page < wv.BaseIndex || int(page) >= len(c.List) {
		return false
	}

	return c.List[page].HasVoice()
}

// GetVoiceUrl will give you a random voice from available voices.
func (c *UrbanCollection) GetVoiceUrls(page uint8) []string {
	if page < wv.BaseIndex || int(page) >= len(c.List) {
		return nil
	}
	return c.List[page].GetVoiceUrls()
}

// Found will give you a boolean indicating whether the word
// is found or not.
func (c *UrbanCollection) Found() bool {
	return len(c.List) != wv.BaseIndex
}

//---------------------------------------------------------

func (d *UrbanDictionary) HasVoice() bool {
	if d.Voices == nil || len(d.Voices) == wv.BaseIndex {
		return false
	}

	return true
}

func (d *UrbanDictionary) GetVoiceUrls() []string {
	return d.Voices
}

//---------------------------------------------------------
