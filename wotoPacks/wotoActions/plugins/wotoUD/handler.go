package wotoUD

import (
	"fmt"
	"log"
	"reflect"
	"strconv"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/pTools"
	wq "github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/wotoQuery"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoMD"
	ws "github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues/tgMessages"
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

		sim, simErr := GetSimilarWord(text)
		simOK := simErr == nil && sim != nil && len(sim) != wv.BaseIndex

		if !simOK {
			finalText = wotoMD.GetNormal(tmpStr).ToString()
		} else {
			tmp := wotoMD.GetNormal(tmpStr)
			mean := simNotice
			num := wv.BaseOneIndex
			for _, s := range sim {
				if ws.IsEmpty(&s) {
					continue
				}

				mean += strconv.Itoa(num) + numSep + s + wv.N_ESCAPE
				num++
			}
			tmp = tmp.Append(wotoMD.GetItalic(mean))
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
		msg.ReplyMarkup = getButton(text, mId, uId, ud, unique)
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

func getButton(word string, id int, userId int64,
	c *UrbanCollection, unique string) *tg.InlineKeyboardMarkup {

	origin := getNewOrigin(word, id, userId, unique)
	origin.setCollection(c)
	pre := getPreviousUdQuery(origin)
	next := getNextUdQuery(origin)
	preID := wq.SetInMap(pre)
	nextID := wq.SetInMap(next)
	q1 := wq.GetQuery(wq.UdQueryPlugin, preID, unique)
	q2 := wq.GetQuery(wq.UdQueryPlugin, nextID, unique)
	b01 := tg.NewInlineKeyboardButtonData(preText, q1.ToString())
	b02 := tg.NewInlineKeyboardButtonData(nextText, q2.ToString())
	buttons := tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			b01,
			b02,
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

	if query.From.ID != origin.uId {
		config := tg.NewCallbackWithAlert(query.ID, notAllowed)
		_, err := api.Request(config)
		if err != nil {
			log.Println(err)
		}
		return
	}

	//log.Println("B :", origin.currentPage)

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
