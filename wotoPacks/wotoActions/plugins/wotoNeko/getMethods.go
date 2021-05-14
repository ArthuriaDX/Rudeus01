// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoNeko

import (
	"strings"

	ws "github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
)

func (n *NekoBase) IsPhoto() bool {
	tmp := strings.ToLower(n.Url)
	b1 := strings.HasSuffix(tmp, wv.Jpg)
	b2 := strings.HasSuffix(tmp, wv.Jpeg)
	b3 := strings.HasSuffix(tmp, wv.Png)
	return b1 || b2 || b3
}

func (n *NekoBase) IsGif() bool {
	tmp := strings.ToLower(n.Url)
	return strings.HasSuffix(tmp, wv.Gif)
}

func (n *NekoBase) IsText() bool {
	b1 := ws.IsEmpty(&n.Url)
	b2 := !ws.IsEmpty(&n.Why)
	b3 := !ws.IsEmpty(&n.Name)
	b4 := !ws.IsEmpty(&n.Cat)
	b5 := !ws.IsEmpty(&n.Fact)
	b6 := !ws.IsEmpty(&n.Owo)
	return b1 && (b2 || b3 || b4 || b5 || b6)
}

func (n *NekoBase) GetText() string {
	if !n.IsText() {
		return wv.EMPTY
	}

	if !ws.IsEmpty(&n.Why) {
		return n.Why
	}
	if !ws.IsEmpty(&n.Name) {
		return n.Name
	}
	if !ws.IsEmpty(&n.Cat) {
		return n.Cat
	}
	if !ws.IsEmpty(&n.Fact) {
		return n.Fact
	}
	if !ws.IsEmpty(&n.Owo) {
		return n.Owo
	}

	return wv.EMPTY
}
