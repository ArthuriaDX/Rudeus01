// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package botCommands

// non-sudo commands
const (
	TMORSe_CMD    = "tmorse"    // wotoMorse plugin
	FMORSE_CMD    = "fmorse"    // wotoMorse plugin
	TR_CMD        = "tr"        // wotoTranslate plugin
	TL_CMD        = "tl"        // wotoTranslate plugin
	TRANSLATE_CMD = "translate" // wotoTranslate plugin
	PAT_CMD       = "pat"       // wotoPat plugin
	PATTU_CMD     = "pattu"     // wotoPat plugin
)

var cmdList map[string]CmdHandler

func cmdListInit() {
	if cmdList != nil {
		return
	}
	// log.Println("In Init!")
	cmdList = make(map[string]CmdHandler)
	add_morseCmdList()
}

func add_morseCmdList() {
	if cmdList != nil {
		if cmdList[TMORSe_CMD] == nil {
			cmdList[TMORSe_CMD] = toMorse_handler
		}
		if cmdList[FMORSE_CMD] == nil {
			cmdList[FMORSE_CMD] = fromMorse_handler
		}

		if cmdList[TR_CMD] == nil {
			cmdList[TR_CMD] = fromMorse_handler
		}
		if cmdList[TRANSLATE_CMD] == nil {
			cmdList[TRANSLATE_CMD] = fromMorse_handler
		}

		if cmdList[PAT_CMD] == nil {
			cmdList[PAT_CMD] = pat_handler
		}
		if cmdList[PATTU_CMD] == nil {
			cmdList[PATTU_CMD] = pat_handler
		}
	}
}
