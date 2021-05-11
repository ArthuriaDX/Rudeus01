// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoNeko

type sfwNeko string

// sfw endpoints for nekos.life api.
const (
	Tickle   sfwNeko = "tickle"
	Slap     sfwNeko = "slap"
	Poke     sfwNeko = "poke"
	Pat      sfwNeko = "pat"
	Neko     sfwNeko = "neko"
	Meow     sfwNeko = "meow"
	Lizard   sfwNeko = "lizard"
	Kiss     sfwNeko = "kiss"
	Hug      sfwNeko = "hug"
	Fox_Girl sfwNeko = "fox_girl"
	Feed     sfwNeko = "feed"
	Cuddle   sfwNeko = "cuddle"
	// news:
	Kemonomimi sfwNeko = "kemonomimi"
	Holo       sfwNeko = "holo"
	Smug       sfwNeko = "smug"
	Baka       sfwNeko = "baka"
	Woof       sfwNeko = "woof"
	Goose      sfwNeko = "ngif"
	Gecg       sfwNeko = "gecg"
	Avatar     sfwNeko = "avatar"
	Waifu      sfwNeko = "waifu"
	// Wallpaper   sfwNeko = "wallpaper" needs improvements
)

// text constants (Lol, these should be counted as sfw, right? XD)
const (
	WhyText  = "why"
	NameText = "name"
	CatText  = "cat"
	FactText = "fact"
	OwoText  = "owoify"
)

var sfw []sfwNeko

func initSfw() {
	if sfw == nil {
		sfw = []sfwNeko{
			Tickle,
			Slap,
			Poke,
			Pat,
			Neko,
			Meow,
			Lizard,
			Kiss,
			Hug,
			Fox_Girl,
			Feed,
			Cuddle,
			Kemonomimi,
			Holo,
			Smug,
			Baka,
			Woof,
			Goose,
			Gecg,
			Avatar,
			Waifu,
		}
	}
}
