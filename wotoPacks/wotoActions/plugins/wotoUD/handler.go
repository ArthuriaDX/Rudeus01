// Rudeus Telegram Bot Project
// Copyright (C) 2021 wotoTeam, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoUD

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/interfaces"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/pTools"
	wq "github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/wotoQuery"
	wm "github.com/ALiwoto/rudeus01/wotoPacks/wotoMD"
	ws "github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues/tgMessages"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues/tgMessages/tgForbidden"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func UdHandler(message *tg.Message, args pTools.Arg) {
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

	pv := args.HasFlag(PRIVATE_FLAG, PV_FLAG)
	// hasMsg := args.HasFlag(MSG_FLAG, MESSAGE_FLAG)
	reply := message.ReplyToMessage != nil
	del := reply && args.HasFlag(DEL_FLAG, DELETE_FLAG)
	single := args.HasFlag(SINGLE_FLAG)
	text := args.JoinNoneFlags(false)

	if ws.IsEmpty(&text) {
		return
	}

	finalText, ud := defineText(text)
	found := ud.Found()

	if !found {
		tmpStr := fmt.Sprintf(notFound, text)

		sim, simErr := GetSimilarWords(text)
		simOK := simErr == nil && sim != nil && len(sim) != wv.BaseIndex

		if !simOK {
			finalText = wm.GetNormal(tmpStr).ToString()
		} else {
			tmp := wm.GetNormal(tmpStr)
			mean := simNotice
			num := wv.BaseOneIndex
			for _, s := range sim {
				if ws.IsEmpty(&s) {
					continue
				}

				mean += strconv.Itoa(num) + numSep + s + wv.N_ESCAPE
				num++
			}
			tmp = tmp.Append(wm.GetItalic(mean))
			finalText = tmp.ToString()
		}
	}

	var msg tg.MessageConfig

	if pv {
		msg = tg.NewMessage(message.From.ID, finalText)
	} else {
		msg = tg.NewMessage(message.Chat.ID, finalText)
	}

	// !pv => for fixing: Bad Request: replied message not found
	if reply && !pv && found {
		r := message.ReplyToMessage
		if r != nil {
			msg.ReplyToMessageID = r.MessageID
		} else {
			msg.ReplyToMessageID = message.MessageID
		}
	} else {
		// for fixing: Bad Request: replied message not found
		if !pv {
			msg.ReplyToMessageID = message.MessageID
		}
	}

	if !single && found {
		mId := message.MessageID
		uId := message.From.ID
		unique := strconv.FormatInt(int64(message.Date), wv.BaseTen)
		msg.ReplyMarkup = getButtons(text, mId, uId, ud, unique)
	}

	if del && found {
		req := tg.NewDeleteMessage(message.Chat.ID, message.MessageID)
		// don't check error or response, we have
		// more important things to do
		go settings.GetAPI().Request(req)
	}

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

func defineText(word string) (str string, c *UrbanCollection) {
	ud, err := Define(word)
	if err != nil {
		log.Println(err)
		return wv.EMPTY, nil
	}

	return ud.GetDef(wv.BaseIndex), ud
}

// getButton will give you the button(s).
// also it wil do the dirty works for origin.
// it will save the buttons in map.
func getButtons(word string, id int, userId int64,
	c *UrbanCollection, unique string) *tg.InlineKeyboardMarkup {

	origin := getNewOrigin(word, id, userId, unique)
	origin.setCollection(c)
	pre := getPreviousUdQuery(origin)
	next := getNextUdQuery(origin)
	voice := getVoiceUdQuery(origin)
	preID := wq.SetInMap(pre)
	nextID := wq.SetInMap(next)
	voiceID := wq.SetInMap(voice)
	q1 := wq.GetQuery(wq.UdQueryPlugin, preID, unique)
	q2 := wq.GetQuery(wq.UdQueryPlugin, nextID, unique)
	q3 := wq.GetQuery(wq.UdQueryPlugin, voiceID, unique)
	b01 := tg.NewInlineKeyboardButtonData(preText, q1.ToString())
	b02 := tg.NewInlineKeyboardButtonData(nextText, q2.ToString())
	b03 := tg.NewInlineKeyboardButtonData(voiceText, q3.ToString())
	buttons := tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			b01,
			b02,
		),
		tg.NewInlineKeyboardRow(
			b03,
		),
	)
	origin.setButtons(&buttons)

	return &buttons
}

func QUdHanler(query *tg.CallbackQuery, q *wq.QueryBase) {
	data := q.GetDataQuery()
	if data == nil {
		// this the point where we have to show the user:
		// list is not available anymore.
		// NOTE: in new version, we will simply remove the
		// buttons, we won't remove the text itself anymore.
		notAvialable(query)
		return
	}

	udData := &udQuery{}
	if reflect.TypeOf(udData) != reflect.TypeOf(data) {
		return
	}

	udData = data.(*udQuery)
	origin := udData.origin

	settings := appSettings.GetExisting()
	if settings == nil {
		return
	}

	api := settings.GetAPI()
	if api == nil {
		return
	}

	sendAlert := func(text string) {
		config := tg.NewCallbackWithAlert(query.ID, text)
		_, err := api.Request(config)
		if err != nil {
			log.Println(err)
		}
	}

	if query.From.ID != origin.uId && !udData.voice {
		sendAlert(notAllowed)
		return
	}

	if udData.next {
		if origin.currentPage == uint8(len(origin.collection.List)) {
			origin.currentPage = wv.BaseOneIndex
		} else {
			origin.currentPage++
		}
	} else if udData.previous {
		if origin.currentPage == wv.BaseOneIndex {
			origin.currentPage = uint8(len(origin.collection.List))
		} else {
			origin.currentPage--
		}
	} else if udData.voice {
		col := origin.collection
		if col == nil {
			return
		}

		if !col.HasVoice(origin.currentPage) {
			sendAlert(noVoices)
			return
		}

		vUrls := col.GetVoiceUrls(origin.currentPage)
		if vUrls == nil {
			sendAlert(noVoices)
			return
		}

		vText := wm.GetNormal(fmt.Sprintf(voiceCap, origin.word))
		tmpV := wv.EMPTY
		index := wv.BaseOneIndex
		iStr := wv.EMPTY
		var tmpMd, tmpH, tmpH2 interfaces.WMarkDown

		for _, current := range vUrls {
			iStr = strconv.Itoa(index)
			tmpV = wv.N_ESCAPE + iStr + numSep
			tmpMd = wm.GetNormal(tmpV)
			iStr = voiceStr + iStr
			tmpH2 = wm.GetHyperLink(iStr, current)
			tmpH = tmpMd.Append(tmpH2)
			vText = vText.Append(tmpH)
			index++
		}

		vConfig := tg.NewMessage(query.From.ID, vText.ToString())
		vConfig.ParseMode = tg.ModeMarkdownV2
		_, err := api.Send(vConfig)
		if err != nil {
			if strings.Contains(err.Error(), tgForbidden.START_CHAT) {
				sendAlert(startVoice)
			}
			log.Println(err)
		} else {
			var cStr string
			if query.Message.Chat.IsPrivate() {
				cStr = sentPv
			} else {
				cStr = checkPv
			}
			callConfig := tg.NewCallback(query.ID, cStr)
			_, err := api.Request(callConfig)
			if err != nil {
				log.Println(err)
			}
		}
		// NOTE:
		// it seems we can't upload it to telegram,
		// well it doesn't metter actually.
		// instead of sending it as voice, we can send the link.
		// 		note in : 13 May 2021, BY: ALiwoto

		return
	} else {
		return
	}

	//log.Println("B :", origin.currentPage)
	page := origin.currentPage - wv.BaseOneIndex
	text := origin.collection.GetDef(page)
	//log.Println(text)
	chat := query.Message.Chat.ID
	msgId := query.Message.MessageID

	eConfig := tg.NewEditMessageText(chat, msgId, text)
	eConfig.ReplyMarkup = origin.keyboard
	eConfig.ParseMode = tg.ModeMarkdownV2

	_, err := api.Request(eConfig)
	if err != nil {
		log.Println(err)
	}
}

func notAvialable(query *tg.CallbackQuery) {
	settings := appSettings.GetExisting()
	if settings == nil {
		return
	}

	api := settings.GetAPI()
	if api == nil {
		return
	}

	chat := query.Message.Chat.ID
	msg := query.Message.MessageID
	config := tg.NewEditMessageText(chat, msg, query.Message.Text)
	config.Entities = query.Message.Entities
	_, err := api.Request(config)
	if err != nil {
		log.Println(err)
		return
	}
}
