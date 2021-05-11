// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package botCommands

import (
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/pTools"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/wotoMorse"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/wotoNeko"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/wotoPat"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/wotoUD"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func toMorseHandler(message *tg.Message, args pTools.Arg) {
	wotoMorse.ToMorseHandler(message, args)
}

func fromMorseHandler(message *tg.Message, args pTools.Arg) {
	wotoMorse.FromMorseHandler(message, args)
}

func patHandler(message *tg.Message, args pTools.Arg) {
	wotoPat.PatHandler(message, args)
}

func tickleNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNeko.TickleHandler(message, args)
}

func slapNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNeko.SlapHandler(message, args)
}

func pokeNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNeko.PokeHandler(message, args)
}

func nekoNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNeko.NekoHandler(message, args)
}

func meowNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNeko.MeowHandler(message, args)
}

func lizardNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNeko.LizardHandler(message, args)
}

func kissNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNeko.KissHandler(message, args)
}

func hugNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNeko.HugHandler(message, args)
}

func foxNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNeko.FoxHandler(message, args)
}

func feedNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNeko.FeedHandler(message, args)
}

func cuddleNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNeko.CuddleHandler(message, args)
}

func kemonomimiNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNeko.KemonomimiHandler(message, args)
}

func holoNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNeko.HoloHandler(message, args)
}

func smugNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNeko.SmugHandler(message, args)
}

func bakaNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNeko.BakaHandler(message, args)
}

func woofNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNeko.WoofHandler(message, args)
}

func gooseNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNeko.GooseHandler(message, args)
}

func gecgNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNeko.GecgHandler(message, args)
}

func avatarNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNeko.AvatarHandler(message, args)
}

func waifuNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNeko.WaifuHandler(message, args)
}

func whyNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNeko.WhyHandler(message, args)
}

func nameNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNeko.NameHandler(message, args)
}

func catNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNeko.CatHandler(message, args)
}

func factNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNeko.FactHandler(message, args)
}

func owoNekoHandler(message *tg.Message, args pTools.Arg) {
	wotoNeko.OwoHandler(message, args)
}

func udHandler(message *tg.Message, args pTools.Arg) {
	wotoUD.UdHandler(message, args)
}
