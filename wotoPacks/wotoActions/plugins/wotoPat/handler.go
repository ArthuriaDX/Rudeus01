// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoPat

import (
	"fmt"
	"log"
	"time"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	wa "github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/common"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/pTools"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/wotoNeko"
	ws "github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues/tgMessages"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues/tgMessages/tgConst"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// PatHandler is entry handler for a pat command.
func PatHandler(message *tg.Message, args pTools.Arg) {
	if message == nil || args == nil {
		return
	}

	text := wv.EMPTY
	settings := appSettings.GetExisting()
	is_hentai := args.HasFlag(HENTAI_FLAG)
	send_pv := args.HasFlag(PRIVATE_FLAG, PV_FLAG)
	is_reply := message.ReplyToMessage != nil
	is_del := is_reply && args.HasFlag(DEL_FLAG, DELETE_FLAG)
	only_photo := args.HasFlag(PH_FLAG, PHOTO_FLAG)
	hasMsg := args.HasFlag(MSG_FLAG, MESSAGE_FLAG)

	if hasMsg {
		text = args.JoinNoneFlags(false)
	}

	if is_hentai {
		if !settings.IsSudo(message.From.ID) {
			sendHentaiError(message, true, false)
			return
		}
		if !message.Chat.IsPrivate() && !send_pv {
			sendHentaiError(message, false, true)
			return
		}

		patHIntit()
		id, t := getRandomHPat()
		if t == NONE || ws.IsEmpty(&id) {
			sendNoPat(message, is_hentai, send_pv)
			return
		}

		sendPat(message, text, id, t, is_hentai, false)
		return
	}

	chance := getChance(only_photo)
	delFunc := func() {
		if is_del {
			req := tg.NewDeleteMessage(message.Chat.ID, message.MessageID)
			// don't check error or response, we have
			// more important things to do
			go settings.GetAPI().Request(req)
		}
	}

	switch chance {
	case wv.BaseIndex:
		// in this case, we have to use `neko is life` api.
		neko, err := wotoNeko.GetRandomPat()
		if err != nil {
			log.Println(err)
			// return
			// don't return it if you got an error!
			// since we don't know what happened (and it doesn't
			// matter at all), we will send a wotopat.
			patInit()

			id, t := getRandomPat()
			if t == NONE || ws.IsEmpty(&id) {
				sendNoPat(message, is_hentai, send_pv)
				return
			}

			delFunc()
			sendPat(message, text, id, t, send_pv, is_reply)
			return
		}

		delFunc()
		sendNeko(message, text, neko, send_pv, is_reply)
		return
	case wv.BaseOneIndex:
		patInit()

		id, t := getRandomPat()
		if t == NONE || ws.IsEmpty(&id) {
			sendNoPat(message, is_hentai, send_pv)
			return
		}

		delFunc()
		sendPat(message, text, id, t, send_pv, is_reply)
	case wv.BaseTwoIndex:
		// in this case, we have to use `headp.at` api.
		neko, err := wotoNeko.GetHeadPat()
		if err != nil {
			log.Println(err)
			// return
			// don't return it if you got an error!
			// since we don't know what happened (and it doesn't
			// matter at all), we will send a wotopat.
			patInit()

			id, t := getRandomPat()
			if t == NONE || ws.IsEmpty(&id) {
				sendNoPat(message, is_hentai, send_pv)
				return
			}

			delFunc()
			sendPat(message, text, id, t, send_pv, is_reply)
			return
		}

		delFunc()
		sendNeko(message, text, neko, send_pv, is_reply)
		return
	}
}

// getChance will give you the chance number.
// If ph is true, it will return BaseOneIndex.
// values:
//	BaseIndex		=> use neko is life.
//	BaseOneIndex 	=> use local db.
//	BaseTwoIndex 	=> use headp.at api.
func getChance(ph bool) int {
	if ph {
		return wv.BaseOneIndex
	}

	return time.Now().Second() % wv.BaseThreeIndex
}

// PatSHandler is entry handler for a sudo pat command.
func PatSHandler(message *tg.Message, args pTools.Arg) {
	settings := appSettings.GetExisting()
	if settings == nil {
		return
	}

	patCli := settings.GetPatClient()
	if patCli == nil {
		return
	}

	patInit()
	is_add := args.HasFlag(ADD_FLAG)
	is_rm := args.HasFlag(RM_FLAG, REM_FLAG, REMOVE_FLAG) && !is_add
	is_hentai := args.HasFlag(HENTAI_FLAG)
	is_reply := message.ReplyToMessage != nil

	if !is_add && !is_rm {
		return
	}

	var id string

	if is_reply {
		photos := message.ReplyToMessage.Photo
		if photos == nil || len(photos) == wv.BaseIndex {
			return
		}

		id = photos[wv.BaseIndex].FileID
		if ws.IsEmpty(&id) {
			return
		}
	} else {
		id = args.GetNonFlags()[wv.BaseIndex]
		if ws.IsEmpty(&id) {
			return
		}
	}

	if is_add {
		if is_hentai {
			exists := false
			if is_reply {
				exists = patArrayHExists(message.ReplyToMessage.Photo)
			} else {
				exists = patHExists(id)
			}
			if exists {
				sendPatExists(message, id, true)
				return
			}

			result, newM := patCli.AddHPat(id, int32(PHOTO_PAT))
			if result != wa.SUCCESS {
				return
			}

			setNewHPat(&newM)
		} else {
			exists := false
			if is_reply {
				exists = patArrayExists(message.ReplyToMessage.Photo)
			} else {
				exists = patExists(id)
			}
			if exists {
				sendPatExists(message, id, true)
				return
			}

			result, newM := patCli.AddPat(id, int32(PHOTO_PAT))
			if result != wa.SUCCESS {
				return
			}

			setNewPat(&newM)
		}
	}

	if is_rm {
		if is_hentai {
			exists := false
			if is_reply {
				exists = patArrayHExists(message.ReplyToMessage.Photo)
			} else {
				exists = patHExists(id)
			}
			if !exists {
				sendPatExists(message, id, false)
				return
			}

			result, newM := patCli.RemoveHPat(id)
			if result != wa.SUCCESS {
				return
			}

			setNewHPat(&newM)
		} else {
			exists := false
			if is_reply {
				exists = patArrayExists(message.ReplyToMessage.Photo)
			} else {
				exists = patExists(id)
			}
			if !exists {
				sendPatExists(message, id, false)
				return
			}

			result, newM := patCli.RemovePat(id)
			if result != wa.SUCCESS {
				return
			}

			setNewPat(&newM)
		}
	}

	sendPatAddNotice(message, is_add, is_hentai)
}

// sendPat will send the pat (most probably :| )
func sendPat(message *tg.Message, caption, ID string, pType patType, pv, reply bool) {
	if message == nil {
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

	var photo tg.PhotoConfig
	if pv {
		photo = tg.NewPhoto(message.From.ID, tg.FileID(ID))
		if message.From.ID == message.Chat.ID {
			photo.ReplyToMessageID = message.MessageID
		}
	} else {
		photo = tg.NewPhoto(message.Chat.ID, tg.FileID(ID))
		if reply && message.ReplyToMessage != nil {
			photo.ReplyToMessageID = message.ReplyToMessage.MessageID
		} else {
			photo.ReplyToMessageID = message.MessageID
		}
	}

	photo.Caption = caption

	if _, err := api.Send(photo); err != nil {

		log.Println(err)
		tgErr := tgMessages.GetTgError(err)
		if tgErr != nil {
			tgErr.SendRandomErrorMessage(message)
		}
		return
	}
}

func sendNeko(message *tg.Message, caption string, neko *wotoNeko.NekoBase,
	pv, reply bool) {
	if message == nil {
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

		photo.Caption = caption
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

		gif.Caption = caption
		sendIt = gif
	} else {
		log.Println(neko.Url)
		return
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

func sendHentaiError(message *tg.Message, sudoErr, gpErr bool) {
	if message == nil {
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

	var msg tg.MessageConfig
	var text string
	id := message.From.ID

	text = message.From.FirstName

	if sudoErr {
		str := SEND_HENTAI_SUDO_ONLY
		tgConst.ReplaceFirstNameMention(&str, text, id)
		msg = tg.NewMessage(message.Chat.ID, str)
	} else if gpErr {
		str := SEND_HENTAI_PV_ONLY
		tgConst.ReplaceFirstNameMention(&str, text, id)
		msg = tg.NewMessage(message.Chat.ID, str)
	} else {
		return
	}

	msg.ReplyToMessageID = message.MessageID

	msg.ParseMode = tg.ModeMarkdownV2
	if _, err := api.Send(msg); err != nil {

		log.Println(err)
		tgErr := tgMessages.GetTgError(err)
		if tgErr != nil {
			tgErr.SendRandomErrorMessage(message)
		}
		return
	}
}

// sendNoPat will send a warn to the user that there is no pat
// data in the database.
func sendNoPat(message *tg.Message, hentai, pv bool) {
	if message == nil {
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

	var msg tg.MessageConfig
	var text string
	var id int64
	if pv {
		id = message.From.ID
	} else {
		id = message.Chat.ID
		msg.ReplyToMessageID = message.MessageID
	}

	text = message.From.FirstName
	var str string
	if hentai {
		str = NO_HPAT
	} else {
		str = NO_PAT
	}

	tgConst.ReplaceFirstNameMention(&str, text, message.From.ID)
	msg = tg.NewMessage(id, str)

	msg.ParseMode = tg.ModeMarkdownV2
	if _, err := api.Send(msg); err != nil {

		log.Println(err)
		tgErr := tgMessages.GetTgError(err)
		if tgErr != nil {
			tgErr.SendRandomErrorMessage(message)
		}
		return
	}
}

// sendPatExists will send a warn to the user that this pat already
// exists in the database.
func sendPatExists(message *tg.Message, id string, exists bool) {
	if message == nil {
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

	var str string
	if exists {
		str = fmt.Sprintf(PAT_EXISTS, id)
	} else {
		str = fmt.Sprintf(PAT_DOESNT_EXIST, id)
	}
	msg := tg.NewMessage(message.Chat.ID, str)

	msg.ReplyToMessageID = message.MessageID

	if _, err := api.Send(msg); err != nil {
		log.Println(err)
		tgErr := tgMessages.GetTgError(err)
		if tgErr != nil {
			tgErr.SendRandomErrorMessage(message)
		}
		return
	}
}

func sendPatAddNotice(message *tg.Message, added, hentai bool) {
	if message == nil {
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

	var str string
	if added {
		if hentai {
			str = HENTAI_PAT_ADDED
		} else {
			str = PAT_ADDED
		}

	} else {
		if hentai {
			str = HENTAI_PAT_REMOVED
		} else {
			str = PAT_REMOVED
		}

	}
	msg := tg.NewMessage(message.Chat.ID, str)

	msg.ReplyToMessageID = message.MessageID

	if _, err := api.Send(msg); err != nil {
		log.Println(err)
		tgErr := tgMessages.GetTgError(err)
		if tgErr != nil {
			tgErr.SendRandomErrorMessage(message)
		}
		return
	}
}
