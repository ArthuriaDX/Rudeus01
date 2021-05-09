// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoPat

import (
	"fmt"
	"log"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	wa "github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/common"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/pTools"
	ws "github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues/tgMessages"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues/tgMessages/tgConst"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Pat_Handler is entry handler for a pat command.
func Pat_Handler(message *tg.Message, args pTools.Arg) {
	if message == nil || args == nil {
		return
	}

	settings := appSettings.GetExisting()
	is_hentai := args.HasFlag(HENTAI_FLAG)
	send_pv := args.HasFlag(PRIVATE_FLAG, PV_FLAG)
	is_reply := message.ReplyToMessage != nil

	if is_hentai {
		if !settings.IsSudo(message.From.ID) {
			send_HentaiError(message, true, false)
			return
		}
		if !message.Chat.IsPrivate() && !send_pv {
			send_HentaiError(message, false, true)
			return
		}

		patHIntit()
		id, t := getRandomHPat()
		if t == NONE || ws.IsEmpty(&id) {
			send_noPat(message, is_hentai, send_pv)
			return
		}

		sendPat(message, id, t, is_hentai, false)
		return
	}

	patIntit()

	id, t := getRandomPat()
	if t == NONE || ws.IsEmpty(&id) {
		send_noPat(message, is_hentai, send_pv)
		return
	}

	sendPat(message, id, t, send_pv, is_reply)
}

// PatS_Handler is entry handler for a sudo pat command.
func PatS_Handler(message *tg.Message, args pTools.Arg) {
	settings := appSettings.GetExisting()
	if settings == nil {
		return
	}

	patCli := settings.GetPatClient()
	if patCli == nil {
		return
	}

	patIntit()
	is_add := args.HasFlag(ADD_FLAG)
	is_rm := args.HasFlag(RM_FLAG, REM_FLAG, REMOVE_FLAG)
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
			if patHExists(id) {
				send_patExists(message, id, true)
			}

			result, newM := patCli.AddHPat(id, int32(PHOTO_PAT))
			if result != wa.SUCCESS {
				return
			}

			setNewHPat(&newM)
		} else {
			if patExists(id) {
				send_patExists(message, id, true)
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
			if !patHExists(id) {
				send_patExists(message, id, false)
			}

			result, newM := patCli.RemoveHPat(id)
			if result != wa.SUCCESS {
				return
			}

			setNewHPat(&newM)
		} else {
			if !patExists(id) {
				send_patExists(message, id, false)
			}

			result, newM := patCli.RemovePat(id)
			if result != wa.SUCCESS {
				return
			}

			setNewPat(&newM)
		}
	}

	send_patAddNotice(message, is_add, is_hentai)
}

func sendPat(message *tg.Message, ID string, pType patType, pv, reply bool) {
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

	if _, err := api.Send(photo); err != nil {

		log.Println(err)
		tgErr := tgMessages.GetTgError(err)
		if tgErr != nil {
			tgErr.SendRandomErrorMessage(message)
		}
		return
	}
}

func send_HentaiError(message *tg.Message, sudoErr, gpErr bool) {
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

// send_noPat will send a warn to the user that there is no pat
// data in the database.
func send_noPat(message *tg.Message, hentai, pv bool) {
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

// send_patExists will send a warn to the user that this pat already
// exists in the database.
func send_patExists(message *tg.Message, id string, exists bool) {
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

func send_patAddNotice(message *tg.Message, added, hentai bool) {
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
