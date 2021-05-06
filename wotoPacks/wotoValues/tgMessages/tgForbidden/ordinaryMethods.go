package tgForbidden

import (
	"time"

	"github.com/ALiwoto/rudeus01/wotoPacks/appSettings"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoMD"
	wv "github.com/ALiwoto/rudeus01/wotoPacks/wotoValues"
	"github.com/ALiwoto/rudeus01/wotoPacks/wotoValues/tgMessages/tgConst"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (_f *ForbiddenError) SendRandomErrorMessage(message *tg.Message) {

	settings := appSettings.GetExisting()
	if settings == nil {
		return
	}

	switch time.Now().Second() % wv.BaseTwoIndex {
	case wv.BaseIndex:
		if _f.IsPrivateMessageError() {
			str := tgConst.PV_WARN_M01
			tgConst.ReplaceDyn(&str, message)

			api := settings.GetAPI()
			if api == nil {
				return
			}

			msg := tg.NewMessage(message.Chat.ID, str)
			msg.ReplyToMessageID = message.MessageID
			msg.ParseMode = tg.ModeMarkdownV2

			pmMe := wotoMD.GetNormal(tgConst.SEND_MESSAGE_M01).ToString()
			burl := tgConst.BOT_START_M01
			tgConst.ReplaceDynBot(&burl, api)

			buttons := tg.NewInlineKeyboardMarkup(
				tg.NewInlineKeyboardRow(
					tg.NewInlineKeyboardButtonURL(pmMe, burl),
				),
			)
			msg.ReplyMarkup = buttons

			if _, err := api.Send(msg); err != nil {
				settings.SendSudo(err.Error())
			}
			return

		} else {
			goto invalidErr
		}
		//md :=
		//str = strings.ReplaceAll("", "", "")
	case wv.BaseOneIndex:
		if _f.IsPrivateMessageError() {
			str := wotoMD.GetNormal(tgConst.PV_WARN_M02).ToString()
			tgConst.ReplaceDyn(&str, message)

			api := settings.GetAPI()
			if api == nil {
				return
			}

			msg := tg.NewMessage(message.Chat.ID, str)
			msg.ReplyToMessageID = message.MessageID
			msg.ParseMode = tg.ModeMarkdownV2

			pmMe := wotoMD.GetNormal(tgConst.SEND_MESSAGE_M02).ToString()
			burl := tgConst.BOT_START_M01
			tgConst.ReplaceDynBot(&burl, api)

			buttons := tg.NewInlineKeyboardMarkup(
				tg.NewInlineKeyboardRow(
					tg.NewInlineKeyboardButtonURL(pmMe, burl),
				),
			)
			msg.ReplyMarkup = buttons
			if _, err := api.Send(msg); err != nil {
				settings.SendSudo(err.Error())
			}
			return

		} else {
			goto invalidErr
		}
	default:

	}

invalidErr:
	settings.SendSudo(_f._value)
}
