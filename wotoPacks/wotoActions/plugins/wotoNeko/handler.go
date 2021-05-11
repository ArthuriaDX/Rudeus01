// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoNeko

import (
	"log"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/pTools"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues/tgMessages"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func TickleHandler(message *tg.Message, args pTools.Arg) {
	base, err := GetRandomTickle()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func SlapHandler(message *tg.Message, args pTools.Arg) {
	base, err := GetRandomSlap()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func PokeHandler(message *tg.Message, args pTools.Arg) {
	base, err := GetRandomPoke()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func NekoHandler(message *tg.Message, args pTools.Arg) {
	base, err := GetRandomNeko()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func MeowHandler(message *tg.Message, args pTools.Arg) {
	base, err := GetRandomMeow()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func LizardHandler(message *tg.Message, args pTools.Arg) {
	base, err := GetRandomLizard()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func KissHandler(message *tg.Message, args pTools.Arg) {
	base, err := GetRandomKiss()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func HugHandler(message *tg.Message, args pTools.Arg) {
	base, err := GetRandomHug()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func FoxHandler(message *tg.Message, args pTools.Arg) {
	base, err := GetRandomFoxGirl()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func FeedHandler(message *tg.Message, args pTools.Arg) {
	base, err := GetRandomFeed()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func CuddleHandler(message *tg.Message, args pTools.Arg) {
	base, err := GetRandomCuddle()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func KemonomimiHandler(message *tg.Message, args pTools.Arg) {
	base, err := GetRandomKemonomimi()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func HoloHandler(message *tg.Message, args pTools.Arg) {
	base, err := GetRandomHolo()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func SmugHandler(message *tg.Message, args pTools.Arg) {
	base, err := GetRandomSmug()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func BakaHandler(message *tg.Message, args pTools.Arg) {
	base, err := GetRandomBaka()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func WoofHandler(message *tg.Message, args pTools.Arg) {
	base, err := GetRandomWoof()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func GooseHandler(message *tg.Message, args pTools.Arg) {
	base, err := GetRandomGoose()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func GecgHandler(message *tg.Message, args pTools.Arg) {
	base, err := GetRandomGecg()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func AvatarHandler(message *tg.Message, args pTools.Arg) {
	base, err := GetRandomAvatar()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func WaifuHandler(message *tg.Message, args pTools.Arg) {
	base, err := GetRandomWaifu()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func WhyHandler(message *tg.Message, args pTools.Arg) {
	base, err := GetRandomWhy()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func NameHandler(message *tg.Message, args pTools.Arg) {
	base, err := GetRandomName()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func CatHandler(message *tg.Message, args pTools.Arg) {
	base, err := GetRandomCat()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func FactHandler(message *tg.Message, args pTools.Arg) {
	base, err := GetRandomFact()
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func OwoHandler(message *tg.Message, args pTools.Arg) {
	str := args.JoinNoneFlags(false)
	if wotoStrings.IsEmpty(&str) {
		return
	}

	base, err := GetOwo(str)
	if err != nil {
		log.Println(err)
		return
	}
	sendNekoBase(message, base, args)
}

func sendNekoBase(message *tg.Message, neko *NekoBase, args pTools.Arg) {
	if message == nil || neko == nil {
		return
	}

	settings := appSettings.GetExisting()
	if settings == nil {
		return
	}

	api := settings.GetAPI()
	if api == nil {
		return
	}

	pv := args.HasFlag(PRIVATE_FLAG, PV_FLAG)
	hasMsg := args.HasFlag(MSG_FLAG, MESSAGE_FLAG)
	reply := message.ReplyToMessage != nil
	del := reply && args.HasFlag(DEL_FLAG, DELETE_FLAG)
	text := args.JoinNoneFlags(false)

	var sendIt tg.Chattable

	if neko.IsPhoto() {
		var photo tg.PhotoConfig
		if pv {
			photo = tg.NewPhoto(message.From.ID, tg.FileURL(neko.Url))
			if message.From.ID == message.Chat.ID {
				photo.ReplyToMessageID = message.MessageID
			}
		} else {
			photo = tg.NewPhoto(message.Chat.ID, tg.FileURL(neko.Url))
			if reply && message.ReplyToMessage != nil {
				photo.ReplyToMessageID = message.ReplyToMessage.MessageID
			} else {
				photo.ReplyToMessageID = message.MessageID
			}
		}
		if hasMsg {
			photo.Caption = text
		}
		sendIt = photo
	} else if neko.IsGif() {
		var gif tg.DocumentConfig
		if pv {
			gif = tg.NewDocument(message.From.ID, tg.FileURL(neko.Url))
			if message.From.ID == message.Chat.ID {
				gif.ReplyToMessageID = message.MessageID
			}
		} else {
			gif = tg.NewDocument(message.Chat.ID, tg.FileURL(neko.Url))
			if reply && message.ReplyToMessage != nil {
				gif.ReplyToMessageID = message.ReplyToMessage.MessageID
			} else {
				gif.ReplyToMessageID = message.MessageID
			}
		}

		if hasMsg {
			gif.Caption = text
		}
		sendIt = gif
	} else if neko.IsText() {
		var msg tg.MessageConfig
		if pv {
			msg = tg.NewMessage(message.From.ID, neko.GetText())
			if message.From.ID == message.Chat.ID {
				msg.ReplyToMessageID = message.MessageID
			}
		} else {
			msg = tg.NewMessage(message.Chat.ID, neko.GetText())
			if reply && message.ReplyToMessage != nil {
				msg.ReplyToMessageID = message.ReplyToMessage.MessageID
			} else {
				msg.ReplyToMessageID = message.MessageID
			}
		}

		if hasMsg {
			msg.Text += text
		}
		sendIt = msg
	} else {
		log.Println(neko.Url)
		return
	}

	if del {
		req := tg.NewDeleteMessage(message.Chat.ID, message.MessageID)
		// don't check error or response, we have
		// more important things to do
		go settings.GetAPI().Request(req)
	}

	if _, err := api.Send(sendIt); err != nil {

		log.Println(err)
		tgErr := tgMessages.GetTgError(err)
		if tgErr != nil {
			tgErr.SendRandomErrorMessage(message)
		}
		return
	}
}
