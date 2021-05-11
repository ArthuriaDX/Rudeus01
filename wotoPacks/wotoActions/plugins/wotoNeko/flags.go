// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoNeko

// the flags for the wotoNekosLife plugin.
// please notice that all flags should begin with
// prefex `--`, but in the pTools, this prefex will be
// removed.
const (
	PV_FLAG      = "pv"      // won't work if message is replied
	HENTAI_FLAG  = "hentai"  // it works only for sudo people
	PRIVATE_FLAG = "private" // won't work if message is replied
	DEL_FLAG     = "del"     // won't work if message is not replied
	DELETE_FLAG  = "delete"  // won't work if message is not replied
	MSG_FLAG     = "msg"     // is there any messages for caption?
	MESSAGE_FLAG = "message" // is there any messages for caption?
)
