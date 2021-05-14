// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package botCommands

import wn "github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/wotoNeko"

// non-sudo commands
const (
	TMorseCmd    = "tmorse"    // wotoMorse plugin
	FMorseCmd    = "fmorse"    // wotoMorse plugin
	TrCmd        = "tr"        // wotoTranslate plugin
	TlCmd        = "tl"        // wotoTranslate plugin
	TranslateCmd = "translate" // wotoTranslate plugin
	PatCmd       = "pat"       // wotoPat plugin
	PattuCmd     = "pattu"     // wotoPat plugin
	FoxCmd       = "fox"       // wotoNekosLife plugin
	FoxyCmd      = "foxy"      // wotoNekosLife plugin
	udCmd        = "ud"        // wotoNekosLife plugin
	defCmd       = "def"       // wotoNekosLife plugin
	defineCmd    = "define"    // wotoNekosLife plugin
)

var cmdList map[string]CmdHandler

func cmdListInit() {
	if cmdList != nil {
		return
	}

	cmdList = make(map[string]CmdHandler)

	addMorseCmdList()
	addTrCmdList()
	addPatCmdList()
	addNekosLifeCmdList()
	addUdCmdList()
	addWpCmdList()
}

func addMorseCmdList() {
	if cmdList != nil {
		if cmdList[TMorseCmd] == nil {
			cmdList[TMorseCmd] = toMorseHandler
		}
		if cmdList[FMorseCmd] == nil {
			cmdList[FMorseCmd] = fromMorseHandler
		}
	}
}

func addTrCmdList() {
	if cmdList != nil {
		if cmdList[TrCmd] == nil {
			cmdList[TrCmd] = trHandler
		}
		if cmdList[TranslateCmd] == nil {
			cmdList[TranslateCmd] = trHandler
		}
	}
}

func addPatCmdList() {
	if cmdList != nil {
		if cmdList[PatCmd] == nil {
			cmdList[PatCmd] = patHandler
		}
		if cmdList[PattuCmd] == nil {
			cmdList[PattuCmd] = patHandler
		}
	}
}

func addNekosLifeCmdList() {
	//first: sfw
	if cmdList != nil {
		if cmdList[string(wn.Tickle)] == nil {
			cmdList[string(wn.Tickle)] = tickleNekoHandler
		}
		if cmdList[string(wn.Slap)] == nil {
			cmdList[string(wn.Slap)] = slapNekoHandler
		}
		if cmdList[string(wn.Poke)] == nil {
			cmdList[string(wn.Poke)] = pokeNekoHandler
		}
		// don't set pat here, we have wotoPat for handling this.
		//if cmdList[string(wn.Pat)] == nil {
		//	cmdList[string(wn.Pat)] = patHandler
		//}
		if cmdList[string(wn.Neko)] == nil {
			cmdList[string(wn.Neko)] = nekoNekoHandler
		}
		if cmdList[string(wn.Meow)] == nil {
			cmdList[string(wn.Meow)] = meowNekoHandler
		}
		if cmdList[string(wn.Lizard)] == nil {
			cmdList[string(wn.Lizard)] = lizardNekoHandler
		}
		if cmdList[string(wn.Kiss)] == nil {
			cmdList[string(wn.Kiss)] = kissNekoHandler
		}
		if cmdList[string(wn.Hug)] == nil {
			cmdList[string(wn.Hug)] = hugNekoHandler
		}
		if cmdList[string(wn.Fox_Girl)] == nil {
			cmdList[string(wn.Fox_Girl)] = foxNekoHandler
		}
		if cmdList[FoxCmd] == nil {
			cmdList[FoxCmd] = foxNekoHandler
		}
		if cmdList[FoxyCmd] == nil {
			cmdList[FoxyCmd] = foxNekoHandler
		}
		if cmdList[string(wn.Feed)] == nil {
			cmdList[string(wn.Feed)] = feedNekoHandler
		}
		if cmdList[string(wn.Cuddle)] == nil {
			cmdList[string(wn.Cuddle)] = cuddleNekoHandler
		}
		if cmdList[string(wn.Kemonomimi)] == nil {
			cmdList[string(wn.Kemonomimi)] = kemonomimiNekoHandler
		}
		if cmdList[string(wn.Holo)] == nil {
			cmdList[string(wn.Holo)] = holoNekoHandler
		}
		if cmdList[string(wn.Smug)] == nil {
			cmdList[string(wn.Smug)] = smugNekoHandler
		}
		if cmdList[string(wn.Baka)] == nil {
			cmdList[string(wn.Baka)] = bakaNekoHandler
		}
		if cmdList[string(wn.Woof)] == nil {
			cmdList[string(wn.Woof)] = woofNekoHandler
		}
		if cmdList[string(wn.Goose)] == nil {
			cmdList[string(wn.Goose)] = gooseNekoHandler
		}
		if cmdList[string(wn.Gecg)] == nil {
			cmdList[string(wn.Gecg)] = gecgNekoHandler
		}
		if cmdList[string(wn.Avatar)] == nil {
			cmdList[string(wn.Avatar)] = avatarNekoHandler
		}
		if cmdList[string(wn.Waifu)] == nil {
			cmdList[string(wn.Waifu)] = waifuNekoHandler
		}
		if cmdList[wn.WhyText] == nil {
			cmdList[wn.WhyText] = whyNekoHandler
		}
		if cmdList[wn.NameText] == nil {
			cmdList[wn.NameText] = nameNekoHandler
		}
		if cmdList[wn.CatText] == nil {
			cmdList[wn.CatText] = catNekoHandler
		}
		if cmdList[wn.FactText] == nil {
			cmdList[wn.FactText] = factNekoHandler
		}
		if cmdList[wn.OwoText] == nil {
			cmdList[wn.OwoText] = owoNekoHandler
		}
	}
}

func addUdCmdList() {
	if cmdList != nil {
		if cmdList[udCmd] == nil {
			cmdList[udCmd] = udHandler
		}
		if cmdList[defCmd] == nil {
			cmdList[defCmd] = udHandler
		}
		if cmdList[defineCmd] == nil {
			cmdList[defineCmd] = udHandler
		}
	}
}

// TODO
func addWpCmdList() {
	// TODO
}
