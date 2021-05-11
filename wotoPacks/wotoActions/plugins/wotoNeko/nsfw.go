// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoNeko

type nsfwNeko string

// nsfw endpoints for nekos.life api.
const (
	Pussy       nsfwNeko = "pussy"
	NsfwNekoGif nsfwNeko = "nsfwnekogif"
	Lewd        nsfwNeko = "lewd"
	Les         nsfwNeko = "les"
	Kuni        nsfwNeko = "kuni"
	Cum         nsfwNeko = "cum"
	Classic     nsfwNeko = "classic"
	Boobs       nsfwNeko = "boobs"
	Bj          nsfwNeko = "bj"
	Anal        nsfwNeko = "anal"
	NsfwAvatar  nsfwNeko = "nsfwavatar"
	Yuri        nsfwNeko = "yuri"
	Trap        nsfwNeko = "trap"
	Tits        nsfwNeko = "tits"
	SoloG       nsfwNeko = "solog"
	Solo        nsfwNeko = "solo"
	SmallBoobs  nsfwNeko = "smallboobs"
	PWankG      nsfwNeko = "pwankg"
	Pussy_JPG   nsfwNeko = "pussy_jpg"
	Lewd_Kemo   nsfwNeko = "lewd_kemo"
	Lewd_K      nsfwNeko = "lewd_k"
	Keta        nsfwNeko = "keta"
	HoloLewd    nsfwNeko = "hololewd"
	HoloEro     nsfwNeko = "holoero"
	Hentai      nsfwNeko = "hentai"
	Futanari    nsfwNeko = "futanari"
	Femdom      nsfwNeko = "femdom"
	FeetG       nsfwNeko = "feetG"
	EroFeet     nsfwNeko = "erofeet"
	Ero         nsfwNeko = "ero"
	EroK        nsfwNeko = "eroK"
	EroKemo     nsfwNeko = "erokemo"
	EroN        nsfwNeko = "eron"
	EroYuri     nsfwNeko = "eroyuri"
	Cum_JPG     nsfwNeko = "cum_jpg"
	Blowjob     nsfwNeko = "blowjob"
)

var nsfw []nsfwNeko

func initNsfw() {
	if nsfw == nil {
		nsfw = []nsfwNeko{
			Pussy,
			NsfwNekoGif,
			Lewd,
			Les,
			Kuni,
			Cum,
			Classic,
			Boobs,
			Bj,
			Anal,
			NsfwAvatar,
			Yuri,
			Trap,
			Tits,
			SoloG,
			Solo,
			SmallBoobs,
			PWankG,
			Pussy_JPG,
			Lewd_Kemo,
			Lewd_K,
			Keta,
			HoloLewd,
			HoloEro,
			Hentai,
			Futanari,
			Femdom,
			FeetG,
			EroFeet,
			Ero,
			EroK,
			EroKemo,
			EroN,
			EroYuri,
			Cum_JPG,
			Blowjob,
		}
	}
}
