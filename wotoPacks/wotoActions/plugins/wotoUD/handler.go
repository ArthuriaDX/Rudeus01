package wotoUD

import (
	"log"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoActions/plugins/pTools"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoMD"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoSecurity/wotoStrings"
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
	//hasMsg := args.HasFlag(MSG_FLAG, MESSAGE_FLAG)
	reply := message.ReplyToMessage != nil
	//del := reply && args.HasFlag(DEL_FLAG, DELETE_FLAG)
	//single := args.HasFlag(SINGLE_FLAG)
	text := args.JoinNoneFlags(false)
	if wotoStrings.IsEmpty(&text) {
		return
	}

	ud, err := Define(text)
	if err != nil {
		log.Println(err)
		return
	}

	//text = wotoMD.GetNormal(ud.List[0].Def).ToString()
	normal := wotoMD.GetNormal(ud.List[0].Def)
	ex := wotoMD.GetItalic(wv.N_ESCAPE + wv.N_ESCAPE + ud.List[0].Example)
	text = normal.Append(ex).ToString()

	var msg tg.MessageConfig

	if pv {

		//settings.SendSudo(strconv.Itoa(int(message.From.ID)))
		msg = tg.NewMessage(message.From.ID, text)
	} else {
		//settings.SendSudo(strconv.Itoa(int(message.Chat.ID)))
		msg = tg.NewMessage(message.Chat.ID, text)
	}

	// !pv => for fixing: Bad Request: replied message not found
	if reply && !pv {
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
	msg.ReplyMarkup = getButton(0)

	msg.ParseMode = tg.ModeMarkdownV2
	if _, err := api.Send(msg); err != nil {

		log.Println(err)
		tgErr := tgMessages.GetTgError(err)
		tgErr.SendRandomErrorMessage(message)
		return
	}

}

func getButton(page int) *tg.InlineKeyboardMarkup {
	//b01 := tg.NewInlineKeyboardButtonURL(text, burl)
	b01 := tg.NewInlineKeyboardButtonData(previousText, "test2")
	b02 := tg.NewInlineKeyboardButtonData(nextText,
		"12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")
	buttons := tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			b01,
			b02,
		),
	)

	return &buttons
}
